package disablebuiltin

import "database/sql/driver"

type SQL struct{} // want `the methods of "SQL" use pointer receiver and non-pointer receiver.`

func (s SQL) Value() (driver.Value, error) {
	panic("not implemented")
}

func (s *SQL) Scan(src any) error {
	panic("not implemented")
}
