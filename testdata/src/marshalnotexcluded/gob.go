package marshalnotexcluded

type Gob struct{} // want `the methods of "Gob" use pointer receiver and non-pointer receiver.`

func (g Gob) GobEncode() ([]byte, error) {
	panic("not implemented")
}

func (g *Gob) SetData(data []byte) {
	panic("not implemented")
}
