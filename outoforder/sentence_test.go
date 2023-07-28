package outoforder_test

import (
	"adsd/evt"
	"adsd/outoforder"
	"strings"

	"testing"
)

func TestHandleEventsInOrder(t *testing.T) {

	want := "the quick brown fox"
	sentence := outoforder.Sentence{}

	sendTextAsEventsTo(want, &sentence)

	got := sentence.Text()

	if got != want {
		t.Fatalf("Wanted %q but got %q", want, got)
	}

}

func sendTextAsEventsTo(text string, aggregate evt.Aggregate) {
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

	handlerMap := aggregate.EventHandlers()

	for _, event := range events {
		if handler, ok := handlerMap[event.Type]; ok {
			handler(event)
		}
	}
}
