package builtinmethods

type Text struct{}

func (t Text) MarshalText() ([]byte, error) {
	panic("not implemented")
}

func (t *Text) UnmarshalText(b []byte) error {
	panic("not implemented")
}
