package evt

const TextAppendedType = "TextAppended"

type TextAppended struct {
	Text string
}

func (a TextAppended) PayloadType() string {
	return TextAppendedType
}
