package builtinmethods

type Gob struct{}

func (g Gob) GetData() []byte {
	panic("not implemented")
}

func (g *Gob) GobDecode(data []byte) error {
	panic("not implemented")
}
