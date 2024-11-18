package disablebuiltin

type Text struct{} // want `the methods of "Text" use pointer receiver and non-pointer receiver.`

func (t Text) MarshalText() ([]byte, error) {
	panic("not implemented")
}

func (t *Text) UnmarshalText(b []byte) error {
	panic("not implemented")
}
