package main

import (
	"fmt"
	"log"
)

/*
	Реализовать паттерн «состояние».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
*/

type State interface {
	summonHorse() error
	revokeHorse() error
	dealDamage(target string) error
}

type Character struct {
	idle   State
	ride   State
	attack State

	currentState State

	horse  bool
	target string
}

func newCharacter(target string, horse bool) *Character {
	c := &Character{
		target: target,
		horse:  horse,
	}
	idleState := &IdleState{
		character: c,
	}
	rideState := &RideState{
		character: c,
	}
	attackState := &AttackState{
		character: c,
	}
	c.setState(idleState)
	c.idle = idleState
	c.ride = rideState
	c.attack = attackState
	return c
}

func (c *Character) summonHorse() error {
	return c.currentState.summonHorse()
}

func (c *Character) revokeHorse() error {
	return c.currentState.revokeHorse()
}
func (c *Character) dealDamage(target string) error {
	return c.currentState.dealDamage(target)
}

func (c *Character) setState(s State) {
	c.currentState = s
}

type IdleState struct {
	character *Character
}

func (i *IdleState) summonHorse() error {
	fmt.Println("You summon horse")
	i.character.horse = true
	i.character.setState(i.character.ride)
	return nil
}

func (i *IdleState) revokeHorse() error {
	return fmt.Errorf("No horse")
}
func (i *IdleState) dealDamage(target string) error {
	return fmt.Errorf("No target")
}

type RideState struct {
	character *Character
}

func (r *RideState) summonHorse() error {
	return fmt.Errorf("Horse already summon")
}

func (r *RideState) revokeHorse() error {
	fmt.Println("You revoke horse")
	r.character.setState(r.character.idle)
	return nil
}
func (r *RideState) dealDamage(target string) error {
	return fmt.Errorf("No target")
}

type AttackState struct {
	character *Character
}

func (a *AttackState) summonHorse() error {
	return fmt.Errorf("Not possible in this state")
}

func (a *AttackState) revokeHorse() error {
	return fmt.Errorf("Not possible in this state")
}
func (a *AttackState) dealDamage(target string) error {
	fmt.Printf("You deal %s 10 damage\n", target)
	a.character.setState(a.character.idle)
	return nil
}

func main() {
	character := newCharacter("enemy", false)

	err := character.summonHorse()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = character.revokeHorse()
	if err != nil {
		log.Fatalf(err.Error())
	}
	character.setState(character.attack)
	err = character.dealDamage("enemy")
	if err != nil {
		log.Fatalf(err.Error())
	}
}

/*
	Состояние позволяет объектам менять поведение в зависимости от своего состояния
	Извне создается впечатление что изменился класс объекта
	Применимость:
	- когда есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния,
	причем типов состояний много и их код часто меняется
	- когда код класса содержит множество больших, похожих друг на друга условных операторов,
	которые выбирают поведение в зависимости от текущих значений полей
	- когда используется машина состояний, построенная на условных операторах,
	но происходит дублирование кода для похожих состояний и переходов
	Плюсы:
	- избавляет от множества больших условных операторов машины состояний
	- концентрирует в одном месте код, связанный с определенным состоянием
	- упрощает код контекста
	Минусы:
	- может неоправданно иусложнить код, если состояний мало и они редко меняются
*/
