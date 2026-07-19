package planner

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type TaskState string

const (
	TaskPending    TaskState = "pending"
	TaskInProgress TaskState = "in_progress"
	TaskCompleted  TaskState = "completed"
	TaskFailed     TaskState = "failed"
	TaskBlocked    TaskState = "blocked"
)

type Task struct {
	ID          string
	Description string
	State       TaskState
	DependsOn   []string
	CreatedAt   time.Time
	StartedAt   time.Time
	CompletedAt time.Time
	Result      string
	Error       string
}

type Plan struct {
	ID        string
	Objective string
	Tasks     []Task
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Planner struct {
	mu    sync.Mutex
	plans map[string]*Plan
}

func New() *Planner {
	return &Planner{
		plans: make(map[string]*Plan),
	}
}

func (p *Planner) CreatePlan(objective string, tasks []Task) *Plan {
	p.mu.Lock()
	defer p.mu.Unlock()

	plan := &Plan{
		ID:        fmt.Sprintf("plan_%d", time.Now().UnixNano()),
		Objective: objective,
		Tasks:     tasks,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	p.plans[plan.ID] = plan
	return plan
}

func (p *Planner) GetPlan(id string) (*Plan, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	plan, ok := p.plans[id]
	return plan, ok
}

func (p *Planner) ListPlans() []*Plan {
	p.mu.Lock()
	defer p.mu.Unlock()
	result := make([]*Plan, 0, len(p.plans))
	for _, plan := range p.plans {
		result = append(result, plan)
	}
	return result
}

func (p *Planner) UpdateTaskState(planID, taskID string, state TaskState, result, errMsg string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	plan, ok := p.plans[planID]
	if !ok {
		return fmt.Errorf("plan not found: %s", planID)
	}

	for i, task := range plan.Tasks {
		if task.ID != taskID {
			continue
		}
		plan.Tasks[i].State = state
		plan.Tasks[i].Result = result
		plan.Tasks[i].Error = errMsg
		plan.UpdatedAt = time.Now()

		switch state {
		case TaskInProgress:
			plan.Tasks[i].StartedAt = time.Now()
		case TaskCompleted, TaskFailed:
			plan.Tasks[i].CompletedAt = time.Now()
		}

		return nil
	}

	return fmt.Errorf("task not found: %s", taskID)
}

func (p *Planner) GetNextTasks(planID string) []Task {
	p.mu.Lock()
	defer p.mu.Unlock()

	plan, ok := p.plans[planID]
	if !ok {
		return nil
	}

	completed := make(map[string]bool)
	for _, task := range plan.Tasks {
		if task.State == TaskCompleted {
			completed[task.ID] = true
		}
	}

	var ready []Task
	for _, task := range plan.Tasks {
		if task.State != TaskPending {
			continue
		}
		depsMet := true
		for _, dep := range task.DependsOn {
			if !completed[dep] {
				depsMet = false
				break
			}
		}
		if depsMet {
			ready = append(ready, task)
		}
	}

	return ready
}

func (p *Planner) PlanSummary(planID string) string {
	plan, ok := p.GetPlan(planID)
	if !ok {
		return "Plan not found."
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("Plan: %s\n", plan.Objective))
	b.WriteString(strings.Repeat("=", 40) + "\n")

	for _, task := range plan.Tasks {
		icon := "o"
		switch task.State {
		case TaskInProgress:
			icon = "*"
		case TaskCompleted:
			icon = "+"
		case TaskFailed:
			icon = "x"
		case TaskBlocked:
			icon = "!"
		}

		deps := ""
		if len(task.DependsOn) > 0 {
			deps = " [depends: " + strings.Join(task.DependsOn, ", ") + "]"
		}

		b.WriteString(fmt.Sprintf("  %s %s%s\n", icon, task.Description, deps))
	}

	return b.String()
}
