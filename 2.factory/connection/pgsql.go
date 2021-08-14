package connection

import (
	"fmt"
	"time"
)

type Postgres struct {
}

func (p *Postgres) Connect() error {
	fmt.Println("Conectado a Postgres")
	return nil
}

func (p *Postgres) GetNow() (s *time.Time, e error) {
	fmt.Println("Obteniendo fecha de Postgres")
	t := time.Now()
	s = &t
	e = nil
	return
}

func (p *Postgres) Close() (e error) {
	fmt.Println("Cerrando la conexion Postgres")
	return
}
