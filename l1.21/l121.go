package main

import "fmt"

//целевой интерфейс
type PowerOutlet interface {
	GetPower() string
}

//адаптируемуя структура
type AmericanLaptopCharger struct {
}

//метод, который не соответствует целевому интерфейсу
func (socket AmericanLaptopCharger) Provide110V() string {
	return "110v"
}

//адаптер
type USToEuroPlugAdapter struct {
	charger *AmericanLaptopCharger
}

//конструктор
func NewUSToEuroPlugAdapter(charger *AmericanLaptopCharger) *USToEuroPlugAdapter {
	return &USToEuroPlugAdapter{charger: charger}
}

//метод целевого интерфейса
func (a *USToEuroPlugAdapter) GetPower() string {
	return a.charger.Provide110V()
}

//работает с целевым интерфейсом
func ClientFunction(outlet PowerOutlet) {
	fmt.Println("Получено напряжение:", outlet.GetPower())
}

func main() {
	usCharger := &AmericanLaptopCharger{}
	// ClientFunction(usCharger) //не скомпилируется
	fmt.Println("Использован адаптер")
	adapter := NewUSToEuroPlugAdapter(usCharger)
	ClientFunction(adapter)
}
