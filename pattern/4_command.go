package main

import "fmt"

/*
	Реализовать паттерн «команда».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
*/

// интерфейс команды
type Command interface {
	execute()
}

// интерфейс принимающего
type Device interface {
	on()
	off()
}

// отправитель
type Button struct{
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

// команды
type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

type Tv struct {
	isRunning bool
}

// получатель
func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}
func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func main() {
	tv := &Tv{}

	onCommand := &OnCommand{device: tv}
	offCommand := &OffCommand{device: tv}

	onButton := &Button{command: onCommand}
	onButton.press()
	offButton := &Button{command: offCommand}
	offButton.press()
}

/*
	Команда превращает запросы в объекты, позволяя передавать их как аргументы при вызове методов,
	ставить запросы в очередь, логировать их, а также поддерживать отмену операций
	Каждый вызов, отличающийся от других, следует завернуть в собственный класс с единственным методом,
	который и будет осуществлять вызов. Такие объекты называют командами.
	Применимость:
	- когда нужно параметризовать объекты выполняемым действием
	- когда нужно ставить операции в очередь, выполнять их по расписанию или передавать по сети
	- когда нужна операция отмены
	Плюсы:
	- убирает прямую зависимость между объектами, вызывающими операции и объектами, которые их выполняют
	- позволяет реализовать простую отмену и повтор операций
	- позволяет реализовать отложенный запуск операций
	- позволяет собирать сложные команды из простых
	- реализует принцип открытости/закрытости
	Минусы:
	- усложняет код программы изза введения множества доп классов
*/
