package builtinmethods

type JSON struct{}

func (j JSON) GetData() []byte {
	panic("not implemented")
}

func (j *JSON) UnmarshalJSON(b []byte) error {
	panic("not implemented")
}
