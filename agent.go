package agents

import (
	"context"
	"sync"

	"github.com/katallaxie/prompts"
)

// Agent represents an agent that can perform tasks and emit events.
type Agent interface {
	// Task performs a task and emits events.
	Task(ctx context.Context, req *TaskRequest) <-chan *Event
	// Error returns any error that occurred during the agent's operation.
	Error() error
}

// Options represents the options for configuring an agent.
type Options struct {
	// SystemPrompt is the system prompt for the agent.
	SystemPrompt prompts.ChatCompletionMessage
	// Client is the client for the agent.
	Client prompts.Prompt
	// Tools is a list of tools that the agent can use to perform tasks.
	Tools []prompts.Tool
}

// Opt is an option for configuring an agent.
type Opt func(*Options)

var _ Agent = (*AgentImpl)(nil)

// AgentImpl represents an agent implementation.
type AgentImpl struct {
	options *Options
	err     error
	errOnce sync.Once
}

// NewAgent creates a new agent with the given options.
func NewAgent(opts ...Opt) *AgentImpl {
	options := new(Options)

	for _, opt := range opts {
		opt(options)
	}

	a := new(AgentImpl)
	a.options = options

	return a
}

// WithClient sets the client for the agent.
func WithClient(client prompts.Prompt) Opt {
	return func(options *Options) {
		options.Client = client
	}
}

// WithSystemPrompt sets the system prompt for the agent.
func WithSystemPrompt(systemPrompt prompts.ChatCompletionMessage) Opt {
	return func(options *Options) {
		options.SystemPrompt = systemPrompt
	}
}

// WithTools sets the tools for the agent.
func WithTools(tools ...prompts.Tool) Opt {
	return func(options *Options) {
		options.Tools = tools
	}
}

// Task performs a task and emits events.
func (a *AgentImpl) Task(_ context.Context, _ *TaskRequest) <-chan *Event {
	events := make(chan *Event)

	return events
}

// Error returns any error that occurred during the agent's operation.
func (a *AgentImpl) Error() error {
	return a.err
}
