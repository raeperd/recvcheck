package marshalnotexcluded

type YAML struct{} // want `the methods of "YAML" use pointer receiver and non-pointer receiver.`

func (y YAML) MarshalYAML() (any, error) {
	panic("not implemented")
}

func (y *YAML) SetData(data []byte) {
	panic("not implemented")
}
