package builtinmethods

type Node struct{}

type YAML struct{}

func (y YAML) GetData() []byte {
	panic("not implemented")
}

func (y *YAML) UnmarshalYAML(value *Node) error {
	panic("not implemented")
}
