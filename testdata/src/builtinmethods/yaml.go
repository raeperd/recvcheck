package builtinmethods

type Node struct{}

type YAML struct{}

func (j YAML) MarshalYAML() (any, error) {
	panic("not implemented")
}

func (j *YAML) UnmarshalYAML(value *Node) error {
	panic("not implemented")
}
