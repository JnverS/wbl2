package main

import (
	"fmt"
	"math"
)

/*
	Реализовать паттерн «посетитель».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
*/

// интерфейс посетителя
type Visitor interface {
	visitorForSquare(*Square)
	visitForCircle(*Circle)
}

// интерфейс фигуры
type Shape interface {
	getType() string
	accept(Visitor)
}

type Square struct {
	side float64
}

func (s *Square) accept (v Visitor) {
	v.visitorForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

type Circle struct {
	r float64
}

func (c *Circle) accept (v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

// посетитель для расчета площади
type AreaCalculator struct {
	area float64
}

func (a *AreaCalculator) visitorForSquare(s *Square) {
	a.area = s.side * s.side
	fmt.Printf("Area of square with side %.1f = %.1f\n", s.side, a.area)
}

func (a *AreaCalculator) visitForCircle(c *Circle) {
	a.area = math.Pi * c.r * c.r
	fmt.Printf("Area of circle with radius %.1f = %.1f\n", c.r, a.area)
}

// посетитель для расчета периметра
type PerimeterCalculator struct {
	P float64
}

func (p *PerimeterCalculator) visitorForSquare(s *Square) {
	p.P = 4 * s.side
	fmt.Printf("Perimetr of square with side %.1f = %.1f\n", s.side, p.P)
}

func (p *PerimeterCalculator) visitForCircle(c *Circle) {
	p.P = 2 * math.Pi * c.r
	fmt.Printf("Perimetr of circle with radius %.1f = %.1f\n", c.r, p.P)
}

func main() {
	square := &Square{side: 3.5}
	circle := &Circle{r: 4}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)

	perimetrCalculator := &PerimeterCalculator{}

	square.accept(perimetrCalculator)
	circle.accept(perimetrCalculator)
}

/*
	Посетитель позволяет добавлять в программу новые операции, не изменяя классы объектов,
	над которыми эти операции могут выполняться.
	Применяется:
	- когда нужно выполнить одну и ту же операцию над всеми элементами сложной структуры объектов.
	- когда над объектами сложной структуры надо выполнятся не связанные между собой операции,
	но не хочется засорять ими классы
	- когда новое поведение имеет смысл только для некоторых классов из существующей иерархии
	Плюсы:
	- упрощает добавление операций, работающих со сложными структурами
	- объединяет родственные операции в одном классе
	- может накапливать состояние при обходе структуры элементов
	Минусы:
	- не оправдан, если иерархия часто меняется
	- может привести к нарушению инкапсуляции
*/
