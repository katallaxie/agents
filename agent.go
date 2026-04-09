package agents

import (
	"context"
	"sync"

	"github.com/katallaxie/prompts"
	"github.com/katallaxie/prompts/callbacks"
)

// Agent represents an agent that can perform tasks and emit events.
type Agent interface {
	// Do performs a task and emits events.
	Do(ctx context.Context, req *TaskRequest) error
}

// ToolExecutionMode represents the mode of tool execution for the agent.
type ToolsExecutionMode string

const (
	// ToolsExecutionModeParallel indicates that tools should be executed in parallel.
	ToolsExecutionModeParallel ToolsExecutionMode = "parallel"
	// ToolsExecutionModeSequential indicates that tools should be executed sequentially.
	ToolsExecutionModeSequential ToolsExecutionMode = "sequential"
)

// Options represents the options for configuring an agent.
type Options struct {
	// SystemPrompt is the system prompt for the agent.
	SystemPrompt prompts.ChatCompletionMessage
	// Client is the client for the agent.
	Client prompts.Prompt
	// Tools is a list of tools that the agent can use to perform tasks.
	Tools []prompts.Tool
	// ToolsExecutionMode is the mode of tool execution for the agent.
	ToolsExecutionMode ToolsExecutionMode
}

// DefaultOptions returns the default options for an agent.
func DefaultOptions() *Options {
	return &Options{
		SystemPrompt:       prompts.ChatCompletionMessage{},
		Client:             nil,
		Tools:              []prompts.Tool{},
		ToolsExecutionMode: ToolsExecutionModeParallel,
	}
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
	options := DefaultOptions()

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

// Do performs a task and emits events.
func (a *AgentImpl) Do(ctx context.Context, req *TaskRequest) error {
	messages := make(chan prompts.ChatCompletionMessage, 1024)
	messages <- prompts.ChatCompletionMessage{
		Role:    prompts.RoleUser,
		Content: a.options.SystemPrompt.Content,
	}

	for message := range messages {
		prompt := &prompts.ChatCompletionRequest{
			Messages: []prompts.ChatCompletionMessage{message},
			Tools:    a.options.Tools,
		}

		err := a.options.Client.SendStreamCompletionRequest(ctx, prompt, callbacks.Print)
		if err != nil {
			return err
		}
	}

	return nil
}

// Error returns any error that occurred during the agent's operation.
func (a *AgentImpl) Error() error {
	return a.err
}
