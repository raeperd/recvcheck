package disablebuiltin

type JSON struct{} // want `the methods of "JSON" use pointer receiver and non-pointer receiver.`

func (j JSON) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}

func (j *JSON) UnmarshalJSON(b []byte) error {
	panic("not implemented")
}
