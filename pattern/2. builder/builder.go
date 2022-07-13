package main

import "fmt"

// Паттерн Строитель позволяет определить процесс создания сложного продукта. Позволяет использовать один и тот же код, чтобы создавать разные объекты.
// Пример сложного объекта - документ, состоящий из заголовка, введения и заключения. Для создания такого объекта можно использовать паттерн Строитель.

// Преимущества: пошаговое создание объектов, один и тот же код для создания объектов, изоляция сборки объекта от бизнес-логики.
// Недостатки: усложенение кода

type Builder interface {
	MakeHeader(str string)
	MakeBody(str string)
	MakeFooter(str string)
}

//Director заведующий всем
type Director struct {
	builder Builder
}

//ConcreteBuilder конкретезирующий строитель
type ConcreteBuilder struct {
	document *Document
}

func (d *Director) Build() {
	d.builder.MakeHeader("Header")
	d.builder.MakeBody("Body")
	d.builder.MakeFooter("Footer")
}

// MakeHeader заполняет хедер документа
func (b *ConcreteBuilder) MakeHeader(str string) {
	b.document.Content += "<header>" + str + "</header>"
}

// MakeBody заполняет тело документа
func (b *ConcreteBuilder) MakeBody(str string) {
	b.document.Content += "<body>" + str + "</body>"
}

// MakeFooter заполняет футер документа
func (b *ConcreteBuilder) MakeFooter(str string) {
	b.document.Content += "<footer>" + str + "</footer>"
}

type Document struct {
	Content string
}

// Show показывает контент документа
func (p *Document) Show() string {
	return p.Content
}

func main() {
	document := new(Document)
	director := Director{&ConcreteBuilder{document: document}}
	director.Build()
	fmt.Println(document.Show())
}
