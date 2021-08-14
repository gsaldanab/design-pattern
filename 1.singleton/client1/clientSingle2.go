package client1

import "github.com/gsaldanab/design-pattern/1.singleton/singleton2"

func IncrementAge2() {
	p := singleton2.GetInstance()
	p.IncrementAge()
}
