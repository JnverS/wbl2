package pattern

/*
	Реализовать паттерн «строитель».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
*/

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

type IntelBuilder struct {
	MB string
	Cpu string
	Ram string
}

func newIntelBuilder() *IntelBuilder {
	return &IntelBuilder{}
}

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

type AMDBuilder struct {
	MB string
	Cpu string
	Ram string
}

func newAMDBuilder() *AMDBuilder {
	return &AMDBuilder{}
}

func (a *AMDBuilder) setMotherboard() {
	a.MB = "Asus"
}

func (a *AMDBuilder) setCPU() {
	a.Cpu = "Core i7"
}

func (a *AMDBuilder) setRAM() {
	a.Ram = "4 GB"
}

func (a *AMDBuilder) getPC() PC {
	return PC {
		MB: a.MB,
		Cpu: a.Cpu,
		Ram: a.Ram,
	}
}

type PC struct {
	MB string
	Cpu string
	Ram string
}
