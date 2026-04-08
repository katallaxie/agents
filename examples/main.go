package main

import (
	"context"
	"os"

	"github.com/katallaxie/agents"
	"github.com/katallaxie/prompts/perplexity"
)

// This example demonstrates how to create a completion request with a message
// It then sends the request to the API and prints the last completion content.
func main() {
	client := perplexity.New(perplexity.WithApiKey(os.Getenv("PPLX_API_KEY")))
	agent := agents.NewAgent(agents.WithClient(client))

	task := &agents.TaskRequest{
		Name:        "Write a haiku",
		Description: "Write a haiku about the beauty of nature.",
	}

	_ = agent.Task(context.Background(), task)
}
