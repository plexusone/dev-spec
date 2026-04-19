# Service Deployment

Deploy devspec as a service for API or MCP integration.

!!! note "Future Feature"
    Service deployment is planned but not yet implemented.

## Planned Features

### MCP Server

Model Context Protocol server for AI assistant integration:

```bash
# Start MCP server
devspec serve --mcp --port 3000
```

MCP tools exposed:

- `devspec_check` - Detect SDD type
- `devspec_validate` - Validate structure
- `devspec_info` - Get SDD information
- `devspec_rubrics` - Get evaluation rubrics
- `devspec_evaluate` - Full evaluation

### HTTP API

REST API for integration with other tools:

```bash
# Start HTTP server
devspec serve --http --port 8080
```

Endpoints:

```
GET  /api/v1/types              # List SDD types
GET  /api/v1/types/{type}       # Get type info
GET  /api/v1/types/{type}/rubrics
POST /api/v1/check              # Check directory
POST /api/v1/validate           # Validate specs
POST /api/v1/evaluate           # Evaluate specs
```

### Docker Deployment

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o devspec ./cmd/devspec

FROM alpine:latest
COPY --from=builder /app/devspec /usr/local/bin/
EXPOSE 8080
CMD ["devspec", "serve", "--http", "--port", "8080"]
```

```bash
docker build -t devspec-server .
docker run -p 8080:8080 -e ANTHROPIC_API_KEY=xxx devspec-server
```

### Kubernetes

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: devspec
spec:
  replicas: 2
  selector:
    matchLabels:
      app: devspec
  template:
    metadata:
      labels:
        app: devspec
    spec:
      containers:
        - name: devspec
          image: devspec-server:latest
          ports:
            - containerPort: 8080
          env:
            - name: ANTHROPIC_API_KEY
              valueFrom:
                secretKeyRef:
                  name: devspec-secrets
                  key: anthropic-api-key
---
apiVersion: v1
kind: Service
metadata:
  name: devspec
spec:
  selector:
    app: devspec
  ports:
    - port: 80
      targetPort: 8080
```

## Use Cases

### Web UI

Build a web interface for spec evaluation:

1. Upload or paste spec content
2. Select SDD type
3. View evaluation results
4. Get improvement suggestions

### IDE Extensions

VS Code extension using the HTTP API:

1. Detect SDD type on file open
2. Validate on save
3. Show inline warnings
4. Provide quick fixes

### Slack/Teams Bot

Integration for team notifications:

1. Post spec to channel
2. Bot evaluates and responds
3. Team reviews results
