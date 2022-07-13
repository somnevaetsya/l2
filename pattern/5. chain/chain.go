package main

import "fmt"

// Паттерн Chain позволяет избежать зависимость между отправителем и получателем запроса, а также дать возможность обработки запроса нескольким объектам.
// Запросы передаются по цепочке, пока не будет обработан каким-то объектом.
// Цепочка обработчиков, которые передают запрос друг другу и решают, обрабатывать его или нет

// Если запрос не обработан, то он передается дальше по цепочке.
// Если же он обработан, то паттерн сам решает передавать его дальше или нет.
// Если запрос не обработан ни одним обработчиком, то он просто теряется

// Применяется при разных запросах, которые нужно обработать разными способами, выполнение обработчиков в строгом порядке, динамическое задание обработчиков
// Преимущества: уменьшение зависимости между клиентом и обработчиками, принцип единственной обязанности, принцип открытости/закрытости
// Недостатки: запрос может быть никем не обработан

type Handler interface {
	SendRequest(message string) string
}

type FirstHandler struct {
	next Handler
}

func (h *FirstHandler) SendRequest(message string) (result string) {
	if message == "first" {
		result = "Im handler 1"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

type SecondHandler struct {
	next Handler
}

func (h *SecondHandler) SendRequest(message string) (result string) {
	if message == "second" {
		result = "Im handler 2"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

type ThirdHandler struct {
	next Handler
}

func (h *ThirdHandler) SendRequest(message string) (result string) {
	if message == "third" {
		result = "Im handler 3"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

func main() {
	handlers := &FirstHandler{
		next: &SecondHandler{
			next: &ThirdHandler{},
		},
	}

	fmt.Println(handlers.SendRequest("first"))

}
