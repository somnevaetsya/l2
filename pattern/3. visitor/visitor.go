package main

import "fmt"

// Паттерн Visitor позволяет определить операцию для объектов других классов без изменения этих классов
// Используется для множества объектов с разными интерфейсами, над которыми нужно провести операцию
// Нужно добавить одинаковый набор операций без изменения классов

// Преимущества: добавление новых операций, объединение классов
// Недостатки: добавление новых классов, потому что необходимо обновлять посетителя

type Visitor interface {
	VisitHome(h *Home) string
	VisitFriend(f *Friend) string
}

type Place interface {
	Accept(v Visitor) string
}

type House struct {
}

func (house *House) VisitHome(h *Home) string {
	return "Visit home"
}

func (house *House) VisitFriend(f *Friend) string {
	return "Visit friend"
}

type PlacesToVisit struct {
	houses []Place
}

func (c *PlacesToVisit) Add(p Place) {
	c.houses = append(c.houses, p)
}

func (c *PlacesToVisit) Accept(v Visitor) string {
	var result string
	for _, p := range c.houses {
		result += p.Accept(v)
	}
	return result
}

type Friend struct {
}

type Home struct {
}

func (home *Home) Accept(v Visitor) string {
	return v.VisitHome(home)
}

func (friend *Friend) Accept(v Visitor) string {
	return v.VisitFriend(friend)
}

func main() {
	placesToVisit := new(PlacesToVisit)
	placesToVisit.Add(&Home{})
	placesToVisit.Add(&Friend{})
	fmt.Println(placesToVisit.Accept(&House{}))
}
