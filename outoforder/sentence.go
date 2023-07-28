package outoforder

import "adsd/evt"

type Sentence struct {
	text string
}

func (s *Sentence) Text() string {
	return s.text
}

func (s *Sentence) reduceTextAppended(event *evt.Event) {
	p := event.Payload.(evt.TextAppended)
	s.text = s.text + p.Text
}

func (s *Sentence) EventHandlers() map[string]evt.EventHandler {
	return map[string]evt.EventHandler{
		evt.TextAppendedType: s.reduceTextAppended,
	}
}
