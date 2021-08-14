package singleton_test

import (
	"sync"
	"testing"

	"github.com/gsaldanab/design-pattern/1.singleton/client1"
	"github.com/gsaldanab/design-pattern/1.singleton/client2"
	"github.com/gsaldanab/design-pattern/1.singleton/singleton"
	"github.com/gsaldanab/design-pattern/1.singleton/singleton2"
)

func TestSingle1(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			client1.IncrementAge()
		}()
		go func() {
			defer wg.Done()
			client2.IncrementAge()
		}()
	}
	wg.Wait()

	p := singleton.GetInstance()
	if p.GetAge() == 200 {
		t.Log("exito, Edad actual: ", p.GetAge())
	} else {
		t.Fatal("falló, Edad actual: ", p.GetAge())
	}

}

func TestSingle2(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			client1.IncrementAge2()
		}()
		go func() {
			defer wg.Done()
			client2.IncrementAge2()
		}()
	}
	wg.Wait()

	p := singleton2.GetInstance()
	if p.GetAge() == 200 {
		t.Log("exito, Edad actual: ", p.GetAge())
	} else {
		t.Fatal("falló, Edad actual: ", p.GetAge())
	}
}
