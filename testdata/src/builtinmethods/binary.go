package builtinmethods

type Binary struct{}

func (b Binary) MarshalBinary() ([]byte, error) {
	panic("not implemented")
}

func (b *Binary) UnmarshalBinary(data []byte) error {
	panic("not implemented")
}
