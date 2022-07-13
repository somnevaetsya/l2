package main

import (
	"fmt"
	"log"
)

// Паттерн Factory Method применяется для создания объектов с определенным интерфейсом, реализации которого представляются потомкам.
// Абстрактный класс фабрики задает правила создания продуктов для наследующих фабрик.

// Пример: обработка файлов различных расширение. Фабрика будет получать расширение и возвращать объект, с помощью которого мы сможем манипулировать с нашим файлом

// Преимущества: упрощение добавления новых объектов, реализация принципа открытости/закрытости.
// Недостатки: большая иерархия классов, так как для каждого объекта необходимо создать свой подкласс создателя.

type Creator interface {
	CreateProduct(action string) Product
}

type Product interface {
	Use() string
}

type ConcreteCreator struct {
}

func NewCreator() Creator {
	return &ConcreteCreator{}
}

func (c *ConcreteCreator) CreateProduct(action string) Product {
	var product Product

	switch action {
	case "a":
		product = &ProductA{action}
	case "b":
		product = &ProductB{action}
	default:
		log.Fatalln("Unknown Action")
	}
	return product
}

type ProductA struct {
	action string
}

func (p *ProductA) Use() string {
	return p.action
}

type ProductB struct {
	action string
}

func (p *ProductB) Use() string {
	return p.action
}

func main() {
	factory := NewCreator()
	products := []Product{
		factory.CreateProduct("a"),
		factory.CreateProduct("b"),
	}

	for _, product := range products {
		fmt.Println(product.Use())
	}
}
