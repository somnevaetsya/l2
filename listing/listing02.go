package main

import (
	"fmt"
)

// defer - указывает внутренней функции выполнится перевыходом из внешней функции
// выполняются по принципу LIFO - последняя добавленная функция будет первой на выполнение

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	// оператор помещает вызов функции в стек. При забирании из стека переменная проинициализирована нулем, поэтому 0++ = 1
	defer func() {
		x++
	}()
	x = 1
	return x
}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
