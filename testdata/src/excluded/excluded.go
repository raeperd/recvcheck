package excluded

import (
	"encoding/xml"
)

type Text struct{}

func (t Text) MarshalText() ([]byte, error) {
	panic("not implemented")
}

func (t *Text) UnmarshalText(b []byte) error {
	panic("not implemented")
}

type JSON struct{}

func (j JSON) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}

func (j *JSON) UnmarshalJSON(b []byte) error {
	panic("not implemented")
}

type XML struct{}

func (x XML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	panic("not implemented")
}

func (x *XML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	panic("not implemented")
}

type Binary struct{}

func (b Binary) MarshalBinary() ([]byte, error) {
	panic("not implemented")
}

func (b *Binary) UnmarshalBinary(data []byte) error {
	panic("not implemented")
}

type Gob struct{}

func (g Gob) GobEncode() ([]byte, error) {
	panic("not implemented")
}

func (g *Gob) GobDecode(data []byte) error {
	panic("not implemented")
}
