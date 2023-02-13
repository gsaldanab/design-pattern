package decorator2_test

import (
	"fmt"
	"reflect"
	"testing"

	decorator2 "github.com/gsaldanab/design-pattern/5.decorator2"
)

func TestDecorator(t *testing.T) {
	// precio: 15
	pizzaPrincipal := &decorator2.VeggieMania{}
	fmt.Printf("interface of pizzaPrincipal -> %s\n", reflect.TypeOf(pizzaPrincipal))

	// agregar cheeseTopping y aumentar el precio: +10
	pizzaWithCheese := &decorator2.CheeseTopping{
		Pizza: pizzaPrincipal,
	}

	// volver a agregar un topping y aumentar el precio nuevamente: +7
	pizzaWithCheeseAndTomato := decorator2.TomatoTopping{
		Pizza: pizzaWithCheese,
	}
	fmt.Printf("interface of pizzaWithCheeseAndTomato -> %s\n", reflect.TypeOf(pizzaWithCheeseAndTomato))

	fmt.Printf("Price of veggieMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrice())
}
