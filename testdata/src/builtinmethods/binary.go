package builtinmethods

type Binary struct{}

func (b Binary) GetData() []byte {
	panic("not implemented")
}

func (b *Binary) UnmarshalBinary(data []byte) error {
	panic("not implemented")
}
