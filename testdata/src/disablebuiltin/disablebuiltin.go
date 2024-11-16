package disablebuiltin

import (
	"encoding/xml"
)

// see: https://github.com/raeperd/recvcheck/issues/7

type Text struct{} // want `the methods of "Text" use pointer receiver and non-pointer receiver.`

func (t Text) MarshalText() ([]byte, error) {
	panic("not implemented")
}

func (t *Text) UnmarshalText(b []byte) error {
	panic("not implemented")
}

type JSON struct{} // want `the methods of "JSON" use pointer receiver and non-pointer receiver.`

func (j JSON) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}

func (j *JSON) UnmarshalJSON(b []byte) error {
	panic("not implemented")
}

type XML struct{} // want `the methods of "XML" use pointer receiver and non-pointer receiver.`

func (x XML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	panic("not implemented")
}

func (x *XML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	panic("not implemented")
}

type Binary struct{} // want `the methods of "Binary" use pointer receiver and non-pointer receiver.`

func (b Binary) MarshalBinary() ([]byte, error) {
	panic("not implemented")
}

func (b *Binary) UnmarshalBinary(data []byte) error {
	panic("not implemented")
}

type Gob struct{} // want `the methods of "Gob" use pointer receiver and non-pointer receiver.`

func (g Gob) GobEncode() ([]byte, error) {
	panic("not implemented")
}

func (g *Gob) GobDecode(data []byte) error {
	panic("not implemented")
}
