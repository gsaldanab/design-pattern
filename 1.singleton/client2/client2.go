package client2

import "github.com/gsaldanab/design-pattern/1.singleton/singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
