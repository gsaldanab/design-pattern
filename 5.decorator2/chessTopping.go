package decorator2

// segundo decorador
type CheeseTopping struct {
	Pizza IPizza
}

func (c *CheeseTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 10
}
