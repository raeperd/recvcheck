package marshalnotexcluded

type Binary struct{} // want `the methods of "Binary" use pointer receiver and non-pointer receiver.`

func (b Binary) MarshalBinary() ([]byte, error) {
	panic("not implemented")
}

func (b *Binary) SetData(data []byte) {
	panic("not implemented")
}
