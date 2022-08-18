package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
*/

type IChocolate interface {
	setName(name string)
	setWeight(weight int)
	getName() string
	getWeight() int
}

type Chocolate struct {
	name   string
	weight int
}

func (c *Chocolate) setName(name string) {
	c.name = name
}

func (c *Chocolate) setWeight(weight int) {
	c.weight = weight
}

func (c *Chocolate) getName() string {
	return c.name
}

func (c *Chocolate) getWeight() int {
	return c.weight
}

type Mars struct {
	Chocolate
}

func newMars() IChocolate {
	return &Mars{
		Chocolate: Chocolate{
			name:   "Mars",
			weight: 50,
		},
	}
}

type Bounty struct {
	Chocolate
}

func newBounty() IChocolate {
	return &Bounty{
		Chocolate: Chocolate{
			name:   "Bounty",
			weight: 55,
		},
	}
}

// в зависимости какуой тип запросили, возвращаем конкретную струкруру реализующую интерфейс
func getChocolate(chocolateType string) (IChocolate, error) {
	if chocolateType == "mars" {
		return newMars(), nil
	}
	if chocolateType == "bounty" {
		return newBounty(), nil
	}
	return nil, fmt.Errorf("Wrong chocolate type")
}

func main() {
	mars, _ := getChocolate("mars")
	bounty, _ := getChocolate("bounty")

	fmt.Printf("Chocolate: %s, weight: %d\n", mars.getName(), mars.getWeight())
	fmt.Printf("Chocolate: %s, weight: %d\n", bounty.getName(), bounty.getWeight())
}

/*
	Фабричный метод определяет общий интерфейс для создания объектов в суперклассе,
	позволяя подклассам изменять тип создаваемых объектов
	предлагается создавать объекты не напрямую, а через вызов особого фабричного метода
	Применимость:
	- когда заранее известны типи и зависимости объектов, к которым придется работать
	- когда нужно дать возможность пользователю расширять часть фреймворка или библиотеки
	- когда нужно экономить системные ресурсы, повторно используя уже созданные объекты,
	вместо создания новых
	Плюсы:
	- избавляет класс от привязки к конкретным классам продуктов
	- внедряет код производства продуктов в одно место, упрощая поддержку кода
	- упрощает добавление новых продуктов
	- реализует принцип открытости/закрытости
	Минусы:
	- может привести к созданию больших параллельных иерархий классов, так как для каждого продукта
	нужно создать свой подкласс создателя
*/
