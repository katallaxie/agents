package agents

import (
	"github.com/google/uuid"
)

// Event represents an event emitted by an agent.
type Event struct {
	Agent isEventAgent
}

type isEventAgent interface {
	isEventAgent()
}

// NewEvent creates a new Event.
func NewEvent() *Event {
	return &Event{}
}

// Reset is resetting the event.
func (e *Event) Reset() {
	*e = Event{}
}

// GetAgentStarted returns the EventAgentStarted if the event is of that type, or nil otherwise.
func (e *Event) GetAgentStarted() *EventAgentStarted {
	if x, ok := e.Agent.(*EventAgentStarted); ok {
		return x
	}

	return nil
}

// GetAgentStopped returns the EventAgentStopped if the event is of that type, or nil otherwise.
func (e *Event) GetAgentStopped() *EventAgentStopped {
	if x, ok := e.Agent.(*EventAgentStopped); ok {
		return x
	}

	return nil
}

// EventAgentStarted is emitted when an agent starts.
type EventAgentStarted struct {
	ID uuid.UUID
}

func (e *EventAgentStarted) isEventAgent() {}

// NewEventAgentStarted creates a new EventAgentStarted.
func NewEventAgentStarted() *Event {
	return &Event{
		Agent: &EventAgentStarted{},
	}
}

// EventAgentStopped is emitted when an agent stops.
type EventAgentStopped struct {
	ID uuid.UUID
}

func (e *EventAgentStopped) isEventAgent() {}

// NewEventAgentStopped creates a new EventAgentStopped.
func NewEventAgentStopped() *Event {
	return &Event{
		Agent: &EventAgentStopped{},
	}
}
