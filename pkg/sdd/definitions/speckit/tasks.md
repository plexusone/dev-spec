---
file: tasks
display_name: Tasks Document
sections:
  - name: Task List
    required: true
    format: checkbox_list
criteria:
  - id: task_clarity
    weight: 0.30
  - id: plan_coverage
    weight: 0.30
  - id: checkbox_format
    weight: 0.20
  - id: task_ordering
    weight: 0.20
---

# Tasks Document Evaluation

## Criterion: task_clarity (30%)

**Task Clarity** - Tasks are clear and actionable

### GO
Tasks are well-defined:
- Each task is specific and actionable
- Clear definition of done
- Appropriate granularity
- Self-contained tasks

### WARN
Tasks mostly clear:
- Some tasks vague
- Mixed granularity

### NO-GO
Tasks unclear:
- Vague task descriptions
- Cannot determine completion
- Tasks too broad

---

## Criterion: plan_coverage (30%)

**Plan Coverage** - Tasks cover the plan

### GO
Full plan coverage:
- All phases have tasks
- Milestones reflected
- Complete implementation path

### WARN
Partial coverage:
- Most plan items covered
- Some gaps

### NO-GO
Poor coverage:
- Major plan items missing
- Incomplete task list

---

## Criterion: checkbox_format (20%)

**Checkbox Format** - Uses checkbox tracking

### GO
Proper checkbox format:
- All tasks use [ ] or [x] format
- Consistent formatting
- Progress trackable

### WARN
Partial checkbox use:
- Some tasks lack checkboxes
- Inconsistent formatting

### NO-GO
No checkbox format:
- Plain text tasks
- Cannot track progress

---

## Criterion: task_ordering (20%)

**Task Ordering** - Tasks are logically ordered

### GO
Tasks well-ordered:
- Logical sequence
- Dependencies respected
- Grouped appropriately

### WARN
Order mostly logical:
- Some ordering issues
- Groups unclear

### NO-GO
Poor ordering:
- Random order
- Dependencies ignored
