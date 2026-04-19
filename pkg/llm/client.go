// Package llm provides LLM client integration using OmniLLM.
package llm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/plexusone/omnillm"
	"github.com/plexusone/omnillm/provider"
)

var (
	ErrNoProvider = errors.New("no LLM provider configured")
	ErrNoAPIKey   = errors.New("no API key configured")
)

// Config holds LLM client configuration.
type Config struct {
	Provider    string
	Model       string
	APIKey      string
	MaxTokens   int
	Temperature float64
}

// Tool represents a tool that can be called by the LLM.
type Tool interface {
	// Name returns the tool name.
	Name() string
	// Description returns a description of what the tool does.
	Description() string
	// Parameters returns the JSON schema for the tool parameters.
	Parameters() map[string]interface{}
	// Execute runs the tool with the given arguments.
	Execute(ctx context.Context, args json.RawMessage) (string, error)
}

// Client wraps OmniLLM for LLM API calls with tool support.
type Client struct {
	config    Config
	omnillm   *omnillm.ChatClient
	tools     []Tool
	toolMap   map[string]Tool
	maxTurns  int
}

// NewClient creates a new LLM client.
func NewClient(config Config) (*Client, error) {
	// Auto-detect provider from environment if not specified
	if config.Provider == "" {
		config.Provider = detectProvider()
	}

	if config.Provider == "" {
		return nil, ErrNoProvider
	}

	// Auto-detect API key from environment
	if config.APIKey == "" {
		config.APIKey = detectAPIKey(config.Provider)
	}

	if config.APIKey == "" {
		return nil, fmt.Errorf("%w: set %s_API_KEY environment variable", ErrNoAPIKey, config.Provider)
	}

	// Set default model if not specified
	if config.Model == "" {
		config.Model = defaultModel(config.Provider)
	}

	// Set defaults
	if config.MaxTokens == 0 {
		config.MaxTokens = 4096
	}

	// Create OmniLLM client
	providerConfig := omnillm.ProviderConfig{
		Provider: omnillm.ProviderName(config.Provider),
		APIKey:   config.APIKey,
	}

	client, err := omnillm.NewClient(omnillm.ClientConfig{
		Providers: []omnillm.ProviderConfig{providerConfig},
	})
	if err != nil {
		return nil, fmt.Errorf("create omnillm client: %w", err)
	}

	return &Client{
		config:   config,
		omnillm:  client,
		tools:    make([]Tool, 0),
		toolMap:  make(map[string]Tool),
		maxTurns: 10, // Maximum tool call iterations
	}, nil
}

// RegisterTool adds a tool to the client.
func (c *Client) RegisterTool(tool Tool) {
	c.tools = append(c.tools, tool)
	c.toolMap[tool.Name()] = tool
}

// RegisterTools adds multiple tools to the client.
func (c *Client) RegisterTools(tools ...Tool) {
	for _, tool := range tools {
		c.RegisterTool(tool)
	}
}

// Complete sends a completion request to the LLM (simple, no tools).
func (c *Client) Complete(prompt string) (string, error) {
	return c.CompleteWithContext(context.Background(), prompt)
}

// CompleteWithContext sends a completion request with context.
func (c *Client) CompleteWithContext(ctx context.Context, prompt string) (string, error) {
	messages := []provider.Message{
		{Role: provider.RoleUser, Content: prompt},
	}

	maxTokens := c.config.MaxTokens
	req := &provider.ChatCompletionRequest{
		Model:     c.config.Model,
		Messages:  messages,
		MaxTokens: &maxTokens,
	}

	if c.config.Temperature > 0 {
		req.Temperature = &c.config.Temperature
	}

	resp, err := c.omnillm.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("chat completion: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("empty response from LLM")
	}

	return resp.Choices[0].Message.Content, nil
}

// CompleteWithTools sends a completion request with tool support.
// The LLM can call tools and the client handles the tool execution loop.
func (c *Client) CompleteWithTools(ctx context.Context, systemPrompt, userPrompt string) (string, error) {
	messages := []provider.Message{}

	if systemPrompt != "" {
		messages = append(messages, provider.Message{
			Role:    provider.RoleSystem,
			Content: systemPrompt,
		})
	}

	messages = append(messages, provider.Message{
		Role:    provider.RoleUser,
		Content: userPrompt,
	})

	// Convert tools to provider format
	providerTools := c.getProviderTools()

	// Tool execution loop
	for i := 0; i < c.maxTurns; i++ {
		maxTokens := c.config.MaxTokens
		req := &provider.ChatCompletionRequest{
			Model:     c.config.Model,
			Messages:  messages,
			MaxTokens: &maxTokens,
		}

		if c.config.Temperature > 0 {
			req.Temperature = &c.config.Temperature
		}

		if len(providerTools) > 0 {
			req.Tools = providerTools
		}

		resp, err := c.omnillm.CreateChatCompletion(ctx, req)
		if err != nil {
			return "", fmt.Errorf("chat completion: %w", err)
		}

		if len(resp.Choices) == 0 {
			return "", errors.New("empty response from LLM")
		}

		choice := resp.Choices[0]

		// Check if no tool calls - return the response
		if len(choice.Message.ToolCalls) == 0 {
			return choice.Message.Content, nil
		}

		// Add assistant message with tool calls to conversation
		messages = append(messages, provider.Message{
			Role:      provider.RoleAssistant,
			ToolCalls: choice.Message.ToolCalls,
		})

		// Execute each tool and add results
		for _, toolCall := range choice.Message.ToolCalls {
			result, err := c.executeTool(ctx, toolCall)
			if err != nil {
				result = fmt.Sprintf("Error executing tool: %v", err)
			}

			toolCallID := toolCall.ID
			messages = append(messages, provider.Message{
				Role:       provider.RoleTool,
				Content:    result,
				ToolCallID: &toolCallID,
			})
		}
	}

	return "", errors.New("exceeded maximum tool call iterations")
}

// executeTool executes a single tool call.
func (c *Client) executeTool(ctx context.Context, toolCall provider.ToolCall) (string, error) {
	tool, ok := c.toolMap[toolCall.Function.Name]
	if !ok {
		return "", fmt.Errorf("unknown tool: %s", toolCall.Function.Name)
	}

	return tool.Execute(ctx, []byte(toolCall.Function.Arguments))
}

// getProviderTools converts registered tools to provider format.
func (c *Client) getProviderTools() []provider.Tool {
	tools := make([]provider.Tool, 0, len(c.tools))
	for _, tool := range c.tools {
		tools = append(tools, provider.Tool{
			Type: "function",
			Function: provider.ToolSpec{
				Name:        tool.Name(),
				Description: tool.Description(),
				Parameters:  tool.Parameters(),
			},
		})
	}
	return tools
}

// Close closes the client and releases resources.
func (c *Client) Close() error {
	if c.omnillm != nil {
		return c.omnillm.Close()
	}
	return nil
}

// detectProvider detects the LLM provider from environment variables.
func detectProvider() string {
	if os.Getenv("ANTHROPIC_API_KEY") != "" {
		return "anthropic"
	}
	if os.Getenv("OPENAI_API_KEY") != "" {
		return "openai"
	}
	return ""
}

// detectAPIKey gets the API key for a provider from environment.
func detectAPIKey(prov string) string {
	switch prov {
	case "anthropic":
		return os.Getenv("ANTHROPIC_API_KEY")
	case "openai":
		return os.Getenv("OPENAI_API_KEY")
	default:
		return ""
	}
}

// defaultModel returns the default model for a provider.
func defaultModel(prov string) string {
	switch prov {
	case "anthropic":
		return "claude-sonnet-4-20250514"
	case "openai":
		return "gpt-4"
	default:
		return ""
	}
}
