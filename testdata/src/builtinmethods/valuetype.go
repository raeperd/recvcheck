package builtinmethods

// ValueType simulates a value receiver type like time.Time where
// most methods use value receivers but Unmarshal methods must use
// pointer receivers. With Unmarshal excluded by default, only value
// receivers remain and no inconsistency is reported.
// See https://github.com/raeperd/recvcheck/issues/17
type ValueType struct{}

func (v ValueType) GetData() []byte {
	panic("not implemented")
}

func (v ValueType) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}

func (v *ValueType) UnmarshalJSON(b []byte) error {
	panic("not implemented")
}
