package connection

import (
	"fmt"
	"time"
)

type MySql struct {
}

func (m *MySql) Connect() error {
	fmt.Println("Conectado a mysql")
	return nil
}

func (m *MySql) GetNow() (s *time.Time, e error) {
	fmt.Println("Obteniendo fecha de mysql")
	t := time.Now()
	s = &t
	e = nil
	return
}

func (m *MySql) Close() (e error) {
	fmt.Println("Cerrando la conexion mysql")
	return
}
