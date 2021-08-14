package decorator

import "fmt"

type FinDecorator struct {
	creatorItem Creator
}

func NewFinDecorator(ci Creator) Creator {
	return &FinDecorator{ci}
}

func (d *FinDecorator) Create() error {
	defer fmt.Println("finalizando...")
	d.creatorItem.Create()
	return nil
}
