package main

import (
	"fmt"
)

/*
	Реализовать паттерн «стратегия».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
*/

type MoveAlgorithm interface {
	move(c *Character)
}

type Walk struct {
}

func (w *Walk) move(c *Character) {
	fmt.Printf("%s is walking\n", c.name)
}

type Run struct {
}

func (r *Run) move(c *Character) {
	fmt.Printf("%s is runing\n", c.name)
}

type Fly struct {
}

func (f *Fly) move(c *Character) {
	fmt.Printf("%s is flying\n", c.name)
}

type Character struct {
	name          string
	moveAlgorithm MoveAlgorithm
}

func newCharacter(m MoveAlgorithm) *Character {
	return &Character{
		name:          "Rayman",
		moveAlgorithm: m,
	}
}

func (c *Character) setMoveAlgorithm(m MoveAlgorithm) {
	c.moveAlgorithm = m
}

func (c *Character) move() {
	c.moveAlgorithm.move(c)
}

func main() {
	walk := &Walk{}
	character := newCharacter(walk)
	character.move()

	run := &Run{}
	character.setMoveAlgorithm(run)
	character.move()

	fly := &Fly{}
	character.setMoveAlgorithm(fly)
	character.move()
}

/*
	Стратегия определяет семейство схожих алгоритмов и помещает каждый из них в собственный класс,
	после чего алгоритмы можно взаимозаменять прямо во время исполнения программы
	предлагает определить семейство схожих алгоритмов, которые часто изменяются или расширяются,
	и вынести их в собственные классы, называемые стратегиями.
	Применимость:
	- когда нужно использовать разные вариации какого-то алгоритма внутри одного объекта
	- когда есть множество похожих классов, отличающихся только некоторым поведением
	- когда не хочется обнажать детали реализации алгоритмов для других классов
	- когда различные вариации алгоритмов реализованы в виде развесистого условного оператора.
	Каждая ветка такого оператора представляют собой вариацию алгоритма
	Плюсы:
	- горячая замена алгоритмов на лету
	- изолирует код и данные алгоритмов от остальных классов
	- уход от наследования к делегированию
	- реализует принцип открытости/закрытости
	Минусы:
	- усложняет программу за счет доп классов
	- клиент должен знать, в чем разница между стратегиями, чтобы выбрать подходящую
*/
