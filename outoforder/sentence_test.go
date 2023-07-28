package outoforder_test

import (
	"adsd/evt"
	"adsd/outoforder"
	"strings"

	"testing"
)

func TestHandleEventDelivery(t *testing.T) {

	tests := map[string]struct {
		sender eventSender
	}{
		"in order":     {sender: inOrderSend},
		"out of order": {sender: reverseOrderSend},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			want := "the quick brown fox"
			sentence := &outoforder.Sentence{}

			tc.sender(textAsEvents(want), sentence)

			got := sentence.Text()

			if got != want {
				t.Fatalf("Wanted %q but got %q", want, got)
			}

		})
	}

}

// textAsEvents splits a string into append events
func textAsEvents(text string) []*evt.Event {
	words := strings.Split(text, " ")

	events := make([]*evt.Event, len(words))

	for i, word := range words {

		payload := word

		if i > 0 {
			payload = " " + payload
		}

		ta := evt.TextAppended{Text: payload}
		events[i] = evt.NewEvent(ta.PayloadType(), ta)
	}
	return events
}

type eventSender func(events []*evt.Event, aggregate evt.Aggregate)

// inOrderSend sends events to an aggregate in order
func inOrderSend(events []*evt.Event, aggregate evt.Aggregate) {

	handlerMap := aggregate.EventHandlers()

	for _, event := range events {
		if handler, ok := handlerMap[event.Type]; ok {
			handler(event)
		}
	}
}

// reverseOrderSend sends events to an aggregate in reverse order
func reverseOrderSend(events []*evt.Event, aggregate evt.Aggregate) {

	reversed := make([]*evt.Event, len(events))

	for i := range events {
		reversed[len(events)-1-i] = events[i]
	}
	inOrderSend(reversed, aggregate)

}
