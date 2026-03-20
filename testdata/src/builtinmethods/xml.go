package builtinmethods

import "encoding/xml"

type XML struct{}

func (x XML) GetData() []byte {
	panic("not implemented")
}

func (x *XML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	panic("not implemented")
}
