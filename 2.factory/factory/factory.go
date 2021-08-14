package factory

import "github.com/gsaldanab/design-pattern/2.factory/connection"

func Factory(i int) connection.DBConnection {
	switch i {
	case 1:
		return &connection.MySql{}
	case 2:
		return &connection.Postgres{}
	default:
		return nil
	}
}
