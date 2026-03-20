package marshalnotexcluded

import "encoding/xml"

type XML struct{} // want `the methods of "XML" use pointer receiver and non-pointer receiver.`

func (x XML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	panic("not implemented")
}

func (x *XML) SetData(data []byte) {
	panic("not implemented")
}
