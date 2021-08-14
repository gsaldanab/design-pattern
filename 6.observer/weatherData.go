package observer

type WeatherData struct {
	observers []Observer
	temp      float64
	humidity  float64
	pressure  float64
}

func (w *WeatherData) RegisterObserver(o Observer) {
	w.observers = append(w.observers, o)
}

func (w *WeatherData) RemoveObserver(o Observer) {
	w.observers = removeFromSlice(w.observers, o)
}

func (w *WeatherData) NotifyObservers() {
	for _, o := range w.observers {
		o.Update(w.temp, w.humidity, w.pressure)
	}
}

func removeFromSlice(observerList []Observer, o Observer) []Observer {
	length := len(observerList)
	for i, obs := range observerList {
		if obs.GetID() == o.GetID() {
			observerList[length-1], observerList[i] = observerList[i], observerList[length-1]
			return observerList[:length-1]
		}
	}
	return observerList
}

func (w *WeatherData) SetMeasurements(temp float64, humidity float64, pressure float64) {
	w.temp = temp
	w.humidity = humidity
	w.pressure = pressure
	w.NotifyObservers()
}
