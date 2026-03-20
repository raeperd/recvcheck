package marshalnotexcluded

import "encoding/xml"

type JSON struct{} // want `the methods of "JSON" use pointer receiver and non-pointer receiver.`

func (j JSON) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}

func (j *JSON) SetData(data []byte) {
	panic("not implemented")
}

type Text struct{} // want `the methods of "Text" use pointer receiver and non-pointer receiver.`

func (t Text) MarshalText() ([]byte, error) {
	panic("not implemented")
}

func (t *Text) SetData(data []byte) {
	panic("not implemented")
}

type YAML struct{} // want `the methods of "YAML" use pointer receiver and non-pointer receiver.`

func (y YAML) MarshalYAML() (any, error) {
	panic("not implemented")
}

func (y *YAML) SetData(data []byte) {
	panic("not implemented")
}

type XML struct{} // want `the methods of "XML" use pointer receiver and non-pointer receiver.`

func (x XML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	panic("not implemented")
}

func (x *XML) SetData(data []byte) {
	panic("not implemented")
}

type Binary struct{} // want `the methods of "Binary" use pointer receiver and non-pointer receiver.`

func (b Binary) MarshalBinary() ([]byte, error) {
	panic("not implemented")
}

func (b *Binary) SetData(data []byte) {
	panic("not implemented")
}

type Gob struct{} // want `the methods of "Gob" use pointer receiver and non-pointer receiver.`

func (g Gob) GobEncode() ([]byte, error) {
	panic("not implemented")
}

func (g *Gob) SetData(data []byte) {
	panic("not implemented")
}
