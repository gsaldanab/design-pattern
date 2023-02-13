package mysql

import "fmt"

type Mysql struct{}

func New() *Mysql {
	return &Mysql{}
}

func (c *Mysql) Find(id int) string {
	return fmt.Sprintf("data from mysql -> %d", id)
}

func (c *Mysql) Save(data string) error {
	fmt.Printf("save %s to mysql\n", data)
	return nil
}
