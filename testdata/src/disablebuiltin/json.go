package disablebuiltin

type JSON struct{} // want `the methods of "JSON" use pointer receiver and non-pointer receiver.`

func (j JSON) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}

func (j *JSON) UnmarshalJSON(b []byte) error {
	panic("not implemented")
}

// JSONV2 uses the encoding/json/v2 method names.
// jsonEncoder and jsonDecoder stand in for the encoding/json/jsontext types.
type JSONV2 struct{} // want `the methods of "JSONV2" use pointer receiver and non-pointer receiver.`

type (
	jsonEncoder struct{}
	jsonDecoder struct{}
)

func (j JSONV2) MarshalJSONTo(e *jsonEncoder) error {
	panic("not implemented")
}

func (j *JSONV2) UnmarshalJSONFrom(d *jsonDecoder) error {
	panic("not implemented")
}
