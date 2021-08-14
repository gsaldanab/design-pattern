package adapter

import "fmt"

type Bicicleta struct {
	Marca string
	Color string
}

func (b *Bicicleta) Pedalear() {
	fmt.Println("Pedaleando!")
}
