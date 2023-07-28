package outoforder_test

import (
	"adsd/evt"
	"adsd/outoforder"
	"strings"

	"testing"
)

func TestHandleEventsInOrder(t *testing.T) {

	want := "the quick brown fox"
	sentence := &outoforder.Sentence{}

	sendTo(textAsEvents(want), sentence)

	got := sentence.Text()

	if got != want {
		t.Fatalf("Wanted %q but got %q", want, got)
	}

}

// sendTo sends events to an aggregate
func sendTo(events []*evt.Event, aggregate evt.Aggregate) {

	handlerMap := aggregate.EventHandlers()

	for _, event := range events {
		if handler, ok := handlerMap[event.Type]; ok {
			handler(event)
		}
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
