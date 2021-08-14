package singleton2

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type person2 struct {
	age int
}

var singleInstancePerson2 *person2

func GetInstance() *person2 {
	if singleInstancePerson2 == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstancePerson2 == nil {
			fmt.Println("creando")
			singleInstancePerson2 = &person2{}
		} else {
			fmt.Println("ya existe")
		}
	} else {
		fmt.Println("ya existe")
	}
	return singleInstancePerson2
}

func (p *person2) IncrementAge() {
	lock.Lock()
	defer lock.Unlock()
	p.age++
}

func (p *person2) GetAge() int {
	lock.Lock()
	defer lock.Unlock()
	return p.age
}
