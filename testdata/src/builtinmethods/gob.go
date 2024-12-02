package builtinmethods

type Gob struct{}

func (g Gob) GobEncode() ([]byte, error) {
	panic("not implemented")
}

func (g *Gob) GobDecode(data []byte) error {
	panic("not implemented")
}
