package main

import "fmt"

// Паттерн Command позволяет представить запрос в виде объекта, значит команда - объектов. Такие запросы можно ставить в очередь, отменять или возобновлять
// Данный паттерн отделяет объект, который инциирует операцию, от объекта, который знает, как ее выполнить.
// Инициатор знает, как отправить команду

// Преимущества: нет зависимости между вызовом операций и выполнением, отмена и повтор операций, сложные команды из простых
// Недостатки: множество дополнительных классов

type Command interface {
	Execute() string
}

type TurnOnCommand struct {
	receiver *Receiver
}

func (on *TurnOnCommand) Execute() string {
	return on.receiver.TurnOn()
}

type TurnOffCommand struct {
	receiver *Receiver
}

func (off *TurnOffCommand) Execute() string {
	return off.receiver.TurnOff()
}

// Receiver implementation.
type Receiver struct {
}

func (r *Receiver) TurnOn() string {
	return "Turn On"
}

func (r *Receiver) TurnOff() string {
	return "Turn Off"
}

type Invoker struct {
	commands []Command
}

func (i *Invoker) AppendCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) DeleteCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

func (i *Invoker) Execute() string {
	var result string
	for _, command := range i.commands {
		result += command.Execute() + "\n"
	}
	return result
}

func main() {
	invoker := &Invoker{}
	receiver := &Receiver{}

	invoker.AppendCommand(&TurnOnCommand{receiver: receiver})
	invoker.AppendCommand(&TurnOffCommand{receiver: receiver})

	fmt.Println(invoker.Execute())
}
