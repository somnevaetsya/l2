package main

import "fmt"

// Паттерн State позволяет изменять свое поведение в зависимости от внутреннего состояния. Подмена класса обхекта

// Преимущества: избавляет от операторов состояний, упрощение кода
// Недостатки: усложенение кода, если мало состояний

type AlertState interface {
	Alert() string
}

type MobileAlert struct {
	state AlertState
}

func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

func (a *MobileAlert) SetState(state AlertState) {
	a.state = state
}

func CreateAlert() *MobileAlert {
	return &MobileAlert{state: &AlertVibration{}}
}

type AlertVibration struct {
}

func (a *AlertVibration) Alert() string {
	return "Vibration..."
}

type AlertSong struct {
}

func (a *AlertSong) Alert() string {
	return "Song..."
}

func main() {
	alert := CreateAlert()

	fmt.Println(alert.Alert())

	alert.SetState(&AlertSong{})
	fmt.Println(alert.Alert())
}
