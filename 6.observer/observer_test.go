package observer_test

import (
	"testing"

	observer "github.com/gsaldanab/design-pattern/6.observer"
)

func TestObserver(t *testing.T) {
	weatherData := observer.WeatherData{}

	clientSuma := observer.NewClient1(1)
	clientPromedio := observer.NewClient2(2)
	clientMayor := observer.NewClient3(3)

	weatherData.RegisterObserver(&clientSuma)
	weatherData.RegisterObserver(&clientPromedio)
	weatherData.RegisterObserver(&clientMayor)

	t.Log("primer cambio")
	weatherData.SetMeasurements(1.2, 2.2, 3.4)

	t.Log("segundo cambio")
	weatherData.SetMeasurements(1, 6.2, 8.7)

	t.Log("tercer cambio")
	weatherData.SetMeasurements(5, 5, 5)
}
