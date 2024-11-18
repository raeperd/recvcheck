package disablebuiltin

type Node struct{}

type YAML struct{} // want `the methods of "YAML" use pointer receiver and non-pointer receiver.`

func (j YAML) MarshalYAML() (any, error) {
	panic("not implemented")
}

func (j *YAML) UnmarshalYAML(value *Node) error {
	panic("not implemented")
}
