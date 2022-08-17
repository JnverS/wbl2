package main

import (
	"fmt"
)

/*
	Реализовать паттерн «строитель».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
*/

// общий интерфейс для всех строителей
type IBuilder interface {
	setMotherboard()
	setCPU()
	setRAM()
	getPC() PC
}

func getBuilder(builderType string) IBuilder {
	if builderType == "intel" {
		return newIntelBuilder()
	}
	if builderType == "amd" {
		return newAMDBuilder()
	}
	return nil
}

// конкретный строитель для интел
type IntelBuilder struct {
	MB string
	Cpu string
	Ram string
}

func newIntelBuilder() *IntelBuilder {
	return &IntelBuilder{}
}
// реализация методов интерфейса
func (i *IntelBuilder) setMotherboard() {
	i.MB = "Asus"
}

func (i *IntelBuilder) setCPU() {
	i.Cpu = "Core i7"
}

func (i *IntelBuilder) setRAM() {
	i.Ram = "4 GB"
}

func (i *IntelBuilder) getPC() PC {
	return PC {
		MB: i.MB,
		Cpu: i.Cpu,
		Ram: i.Ram,
	}
}

//конкретный строитель для амд
type AMDBuilder struct {
	MB string
	Cpu string
	Ram string
}

func newAMDBuilder() *AMDBuilder {
	return &AMDBuilder{}
}
// реализация методов интерфейса строителя
func (a *AMDBuilder) setMotherboard() {
	a.MB = "Gigabyte"
}

func (a *AMDBuilder) setCPU() {
	a.Cpu = "Ryzen"
}

func (a *AMDBuilder) setRAM() {
	a.Ram = "8 GB"
}

func (a *AMDBuilder) getPC() PC {
	return PC {
		MB: a.MB,
		Cpu: a.Cpu,
		Ram: a.Ram,
	}
}
// струкрура нашго пк
type PC struct {
	MB string
	Cpu string
	Ram string
}
// директор для управления строителями
type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}
// назначаем директору конкретного строителя
func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildPC() PC {
	d.builder.setMotherboard()
	d.builder.setCPU()
	d.builder.setRAM()
	return d.builder.getPC()
}

func main() {
	intelBuilder := getBuilder("intel")
	amdBuilder := getBuilder("amd")

	director := newDirector(intelBuilder)
	intelPC := director.buildPC()

	fmt.Printf("Intel PC MB Type: %s\n", intelPC.MB)
	fmt.Printf("Intel PC CPU Type: %s\n", intelPC.Cpu)
	fmt.Printf("Intel PC RAM Type: %s\n", intelPC.Ram)

	director.setBuilder(amdBuilder)
	amdPC := director.buildPC()

	fmt.Printf("AMD PC MB Type: %s\n", amdPC.MB)
	fmt.Printf("AMD PC CPU Type: %s\n", amdPC.Cpu)
	fmt.Printf("AMD PC RAM Type: %s\n", amdPC.Ram)
}

/*
	Строитель отделяет конструирование сложного объекта от его представления,
	так что один и тот же процесс строительства может создать разные представления.
	Применяется:
	- когда нужно избавиться от "телескопического конструктора"
	- когда нужно создавать разные представления какого-то объекта
	- когда нужно собирать сложные составные объекты
	Плюсы:
	- позволяет создавать объекты пошагово
	- позволяет использовать один и тот же код для создания различных продуктов
	- изолирует сложный код сборки продукта от его основной бизнес-логики
	Минусы:
	- усложняет код программы изза доп классов
	- привязка к конкретным классам строителей, тк в интерфейсе директора не может быть метода получения результата
*/
