package main

import (
	"fmt"
	"math/rand"
	"time"
)

// функция получает значения int, далее с рандомной задержкой асинхронно пишет их в канал и возвращает этот канал. после записи закрывает канал
func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

// функция принимает на вход два канала, асинхронно считывает из них в свой канал выхода, возвращает в канал
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			// получение по закрытом канаду дает нулевое значение после того, как все элементы в канале получены
			// поэтому сначала будет вывод
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}

}
