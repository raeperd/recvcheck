package disablebuiltin

import "encoding/xml"

type XML struct{} // want `the methods of "XML" use pointer receiver and non-pointer receiver.`

func (x XML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	panic("not implemented")
}

func (x *XML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	panic("not implemented")
}
