package exclusions

import "database/sql/driver"

type SQL struct{}

func (s SQL) Value() (driver.Value, error) {
	panic("not implemented")
}

func (s *SQL) Scan(src any) error {
	panic("not implemented")
}
