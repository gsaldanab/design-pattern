package singleton

import "sync"

var (
	p     *person
	once  sync.Once
	mutex sync.Mutex
)

func GetInstance() *person {
	once.Do(func() {
		p = &person{}
	})

	return p
}

type person struct {
	name string
	age  int
}

func (p *person) SetName(n string) {
	p.name = n
}

func (p *person) GetName() string {
	return p.name
}

func (p *person) IncrementAge() {
	mutex.Lock()
	defer mutex.Unlock()
	p.age++
}

func (p *person) GetAge() int {
	mutex.Lock()
	defer mutex.Unlock()
	return p.age
}
