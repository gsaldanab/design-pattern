package decorator2

// componente concreto
type VeggieMania struct {
}

func (p *VeggieMania) GetPrice() int {
	return 15
}
