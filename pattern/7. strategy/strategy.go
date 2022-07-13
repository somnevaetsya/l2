package main

import "fmt"

// Паттерн Strategy определяет набор алгоритмов, схожих по роду деятельности, инкапсулирует их в отдельный класс и делает их подменяемыми.
// Позволяет подменять алгоритмы без участия клиентов, которые используют эти алгоритмы.

// Преищущества: изоляция алгоритмов, уход от наследования к делегированию, реализация принципа открытости/закрытости, смена алгоритмов
// Недостатки: дополнительные классы, разница между стратегиями

type Strategy interface {
	Implementation() string
}

type StrategyA struct {
}

func (s *StrategyA) Implementation() string {
	return "Strategy A"
}

type StrategyB struct {
}

func (s *StrategyB) Implementation() string {
	return "Strategy B"
}

type Context struct {
	str Strategy
}

func (c *Context) SetStrategy(s Strategy) {
	c.str = s
}

func (c *Context) UseStrategy() string {
	return c.str.Implementation()
}

func main() {
	ctx := new(Context)

	ctx.SetStrategy(&StrategyA{})
	fmt.Println(ctx.UseStrategy())

	ctx.SetStrategy(&StrategyB{})
	fmt.Println(ctx.UseStrategy())
}
