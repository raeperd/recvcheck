package disablebuiltin

type Binary struct{} // want `the methods of "Binary" use pointer receiver and non-pointer receiver.`

func (b Binary) MarshalBinary() ([]byte, error) {
	panic("not implemented")
}

func (b *Binary) UnmarshalBinary(data []byte) error {
	panic("not implemented")
}
