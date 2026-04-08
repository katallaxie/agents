package agents

import (
	"github.com/google/uuid"
)

// Event represents an event emitted by an agent.
type Event struct {
	Agent isEvent_Agent
}

type isEvent_Agent interface {
	isEvent_Agent()
}

// NewEvent creates a new Event.
func NewEvent() *Event {
	return &Event{}
}

// Reset is resetting the event.
func (e *Event) Reset() {
	*e = Event{}
}

// GetAgentStarted returns the Event_AgentStarted if the event is of that type, or nil otherwise.
func (e *Event) GetAgentStarted() *Event_AgentStarted {
	if x, ok := e.Agent.(*Event_AgentStarted); ok {
		return x
	}

	return nil
}

// GetAgentStopped returns the Event_AgentStopped if the event is of that type, or nil otherwise.
func (e *Event) GetAgentStopped() *Event_AgentStopped {
	if x, ok := e.Agent.(*Event_AgentStopped); ok {
		return x
	}

	return nil
}

// Event_AgentStarted is emitted when an agent starts.
type Event_AgentStarted struct {
	ID uuid.UUID
}

func (e *Event_AgentStarted) isEvent_Agent() {}

// NewEventAgentStarted creates a new Event_AgentStarted.
func NewEventAgentStarted() *Event {
	return &Event{
		Agent: &Event_AgentStarted{},
	}
}

// Event_AgentStopped is emitted when an agent stops.
type Event_AgentStopped struct {
	ID uuid.UUID
}

func (e *Event_AgentStopped) isEvent_Agent() {}

// NewEventAgentStopped creates a new Event_AgentStopped.
func NewEventAgentStopped() *Event {
	return &Event{
		Agent: &Event_AgentStopped{},
	}
}
