package agents

import (
	"context"

	"github.com/katallaxie/prompts"
)

// Agent represents an agent that can perform tasks and emit events.
type Agent interface {
	// Task performs a task and emits events.
	Task(ctx context.Context, req *TaskRequest) <-chan *Event
}

// Opt is an option for configuring an agent.
type Opt func(*AgentImpl)

var _ Agent = (*AgentImpl)(nil)

// AgentImpl represents an agent implementation.
type AgentImpl struct {
	client prompts.Chat
}

// NewAgent creates a new agent with the given options.
func NewAgent(opts ...Opt) *AgentImpl {
	agent := &AgentImpl{}

	for _, opt := range opts {
		opt(agent)
	}

	return agent
}

// WithClient sets the client for the agent.
func WithClient(client prompts.Chat) Opt {
	return func(agent *AgentImpl) {
		agent.client = client
	}
}

// Task performs a task and emits events.
func (a *AgentImpl) Task(_ context.Context, _ *TaskRequest) <-chan *Event {
	events := make(chan *Event)

	return events
}
