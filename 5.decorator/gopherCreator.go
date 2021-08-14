package decorator

import "fmt"

type GopherCreator struct {
}

func (g *GopherCreator) Create() error {
	fmt.Println("creando...")
	return nil
}
