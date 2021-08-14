package observer

import "fmt"

type Client3 struct {
	id       int64
	temp     float64
	humidity float64
	pressure float64
}

func NewClient3(id int64) Client3 {
	return Client3{id: id}
}

func (c *Client3) Update(temp float64, humidity float64, pressure float64) {
	c.temp = temp
	c.humidity = humidity
	c.pressure = pressure
	c.display()
}

func (c *Client3) GetID() int64 {
	return c.id
}

func (c *Client3) display() {
	fmt.Printf("Cliente3 -> Temperatura: %f, Humedad: %f, Presion: %f| mayor: %f\n", c.temp, c.humidity, c.pressure, c.mayor())
}

func (c *Client3) mayor() float64 {
	if c.temp > c.humidity {
		if c.temp > c.pressure {
			return c.temp
		} else {
			return c.pressure
		}
	} else if c.humidity > c.pressure {
		return c.humidity
	} else {
		return c.pressure
	}
}
