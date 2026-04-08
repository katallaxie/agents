package agents

import (
	"github.com/katallaxie/prompts"
)

// TaskRequest represents a request to perform a task.
type TaskRequest struct {
	// Name is the name of the task.
	Name string
	// Description is a description of the task.
	Description string
	// Messages is a list of messages that provide context for the task.
	Messages []prompts.ChatCompletionMessage
}
