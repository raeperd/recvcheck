package marshalnotexcluded

type Text struct{} // want `the methods of "Text" use pointer receiver and non-pointer receiver.`

func (t Text) MarshalText() ([]byte, error) {
	panic("not implemented")
}

func (t *Text) SetData(data []byte) {
	panic("not implemented")
}
