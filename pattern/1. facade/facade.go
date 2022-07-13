package main

import (
	"fmt"
	"strings"
)

// Фасад - некоторый объект, который содержит в себе набор методов с некоторой сложной подсистемой.
// Фасад включает в себя классы, которые реализуют функционал этой системы, но не скрывает их.
// Клиент не лишается более низкоуровневого доступа к классам подсистемы
// Разбиение на подсистемы позволяет упростит процесс разработки и убрать зависимости между компонентами системы.
// Но использование такой системы становится довольно трудным. Суть фасада - единый интерфейс для управления подсистемами

// В качестве примера можно привести интерфейс автомобиля.
// Современные автомобили имеют унифицированный интерфейс для водителя, под которым скрывается сложная подсистема.
// Благодаря применению навороченной электроники, делающей большую часть работы за водителя, тот может с лёгкостью управлять автомобилем, не задумываясь, как там все работает.

// Преимущество: изолиирует клиентов от компонентов сложной подсистемы
// Недостатки: может стать объектом, который будет привязан ко всем классам программы

func CreateMusicBand() *MusicGroup {
	return &MusicGroup{
		guitar: &Guitarist{},
		bass:   &Bassist{},
		drums:  &Drummer{},
	}
}

type MusicGroup struct {
	guitar *Guitarist
	bass   *Bassist
	drums  *Drummer
}

func (group *MusicGroup) PlaySong() string {
	result := []string{group.drums.playDrums(), group.bass.playBass(), group.guitar.playGuitar()}
	return strings.Join(result, "\n")
}

type Guitarist struct {
}

func (g *Guitarist) playGuitar() string {
	return "Guitarist plays guitar"
}

type Bassist struct {
}

func (b *Bassist) playBass() string {
	return "Bassist plays bass"
}

type Drummer struct {
}

func (d *Drummer) playDrums() string {
	return "Drummer plays drums"
}

func main() {
	group := CreateMusicBand()
	fmt.Println(group.PlaySong())
}
