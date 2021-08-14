package observer

import "fmt"

type Client2 struct {
	id       int64
	temp     float64
	humidity float64
	pressure float64
}

func NewClient2(id int64) Client2 {
	return Client2{id: id}
}

func (c *Client2) Update(temp float64, humidity float64, pressure float64) {
	c.temp = temp
	c.humidity = humidity
	c.pressure = pressure
	c.display()
}

func (c *Client2) GetID() int64 {
	return c.id
}

func (c *Client2) display() {
	fmt.Printf("Cliente2 -> Temperatura: %f, Humedad: %f, Presion: %f| promedio: %f\n", c.temp, c.humidity, c.pressure, c.promedio())
}

func (c *Client2) promedio() float64 {
	return (c.temp + c.humidity + c.pressure) / 3
}
