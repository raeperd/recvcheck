package marshalnotexcluded

type JSON struct{} // want `the methods of "JSON" use pointer receiver and non-pointer receiver.`

func (j JSON) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}

func (j *JSON) SetData(data []byte) {
	panic("not implemented")
}
