package evt

const TextAppendedType = "TextAppended"

type TextAppended struct {
	Text string
}

func NewTextAppended(text string) *TextAppended {
	return &TextAppended{Text: text}
}

func (a TextAppended) PayloadType() string {
	return TextAppendedType
}
