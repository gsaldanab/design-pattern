package observer

import "fmt"

type Client1 struct {
	id       int64
	temp     float64
	humidity float64
	pressure float64
}

func NewClient1(id int64) Client1 {
	return Client1{id: id}
}

func (c *Client1) Update(temp float64, humidity float64, pressure float64) {
	c.temp = temp
	c.humidity = humidity
	c.pressure = pressure
	c.display()
}

func (c *Client1) GetID() int64 {
	return c.id
}

func (c *Client1) display() {
	fmt.Printf("Cliente1 -> Temperatura: %f, Humedad: %f, Presion: %f| suma: %f\n", c.temp, c.humidity, c.pressure, c.sum())
}

func (c *Client1) sum() float64 {
	return c.temp + c.humidity + c.pressure
}
