package agents_test

import (
	"testing"

	"github.com/katallaxie/agents"

	"github.com/stretchr/testify/require"
)

func TestNewAgent(t *testing.T) {
	agent := agents.NewAgent()
	require.NotNil(t, agent)
}
