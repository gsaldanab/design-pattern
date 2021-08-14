package adapter

type AutomovilAdapter struct {
	auto *Automovil
}

func (a *AutomovilAdapter) Mover() {
	a.auto.Encender()
	a.auto.Acelerar()
}
