package test

import (
	"database/sql/driver"
	"encoding/xml"
)

// see: https://github.com/raeperd/recvcheck/issues/7
type TEXT struct{}

func (t TEXT) MarshalText() ([]byte, error) {
	panic("not implemented")
}

func (t *TEXT) UnmarshalText(b []byte) error {
	panic("not implemented")
}

type JSON struct{}

func (po JSON) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}

func (po *JSON) UnmarshalJSON(b []byte) error {
	panic("not implemented")
}

type XML struct{}

func (xm XML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	panic("not implemented")
}

func (xm *XML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	panic("not implemented")
}

type BINARY struct{}

func (b BINARY) MarshalBinary() ([]byte, error) {
	panic("not implemented")
}

func (b *BINARY) UnmarshalBinary(data []byte) error {
	panic("not implemented")
}

type GOB struct{}

func (g GOB) GobEncode() ([]byte, error) {
	panic("not implemented")
}

func (g *GOB) GobDecode(data []byte) error {
	panic("not implemented")
}

type SQL struct{}

func (vp SQL) Value() (driver.Value, error) {
	panic("not implemented")
}

func (vp *SQL) Scan(value interface{}) error {
	panic("not implemented")
}

type SQL2 struct{}

func (vp SQL2) Value() (any, error) {
	panic("not implemented")
}

func (vp *SQL2) Scan(value interface{}) error {
	panic("not implemented")
}

type SQL3 struct{}

func (vp SQL3) Value() (interface{}, error) {
	panic("not implemented")
}

func (vp *SQL3) Scan(value interface{}) error {
	panic("not implemented")
}
