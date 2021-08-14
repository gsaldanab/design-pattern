package decorator

import "fmt"

type InitDecorator struct {
	creatorItem Creator
}

func NewInitDecorator(ci Creator) Creator {
	return &InitDecorator{ci}
}

func (d *InitDecorator) Create() error {
	fmt.Println("iniciando...")
	d.creatorItem.Create()
	return nil
}
