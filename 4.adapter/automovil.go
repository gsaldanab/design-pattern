package adapter

import "fmt"

type Automovil struct {
	Marcar    string
	Modelo    string
	Encendido bool
}

func (a *Automovil) Encender() {
	if a.Encendido {
		fmt.Println("Ya está encendido")
		return
	}
	a.Encendido = true
	fmt.Println("Encendido!")
}

func (a *Automovil) Acelerar() {
	fmt.Println("Acelerando!")
}
