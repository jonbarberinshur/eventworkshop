package evt

type EventPayload interface {
	PayloadType() string
}

// Event model
type Event struct {
	Type    string
	Payload EventPayload
}

func NewEvent(Type string, payload EventPayload) *Event {
	return &Event{Type: Type, Payload: payload}
}

type EventHandler func(event *Event)
