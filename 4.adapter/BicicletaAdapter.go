package adapter

type BicicletaAdapter struct {
	bici *Bicicleta
}

func (b *BicicletaAdapter) Mover() {
	b.bici.Pedalear()
}
