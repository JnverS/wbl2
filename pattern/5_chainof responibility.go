package main

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
*/

type Level interface {
	execute(*Client)
	setNext(Level)
}

type FirstLineSupport struct {
	next Level
}

func (f *FirstLineSupport) execute (c *Client) {
	if c.flDone {
		fmt.Println("First Line Support has already helped the client")
		f.next.execute(c)
		return
	}
	fmt.Println("FirstLine helping the client")
	c.flDone = true
	f.next.execute(c)
}

func (r *FirstLineSupport) setNext(next Level) {
	r.next = next
}

type SecondLineSupport struct {
	next Level
}

func (s *SecondLineSupport) execute (c *Client) {
	if c.slDone {
		fmt.Println("SecondLine Support has already helped the client")
		s.next.execute(c)
		return
	}
	fmt.Println("SecondLine Support helping the client")
	c.slDone = true
	s.next.execute(c)
}

func (d *SecondLineSupport) setNext(next Level) {
	d.next = next
}

type Engineer struct {
    next Level
}

func (e *Engineer) execute(c *Client) {
    if c.engineerDone {
        fmt.Println("Engineer has helped the client")
    }
    fmt.Println("Engineer helping the client")
    c.engineerDone = true
}

func (m *Engineer) setNext(next Level) {
    m.next = next
}

type Client struct {
	name string
	flDone bool
	slDone bool
	engineerDone bool
}


func main() {
	engineer := &Engineer{}

	slSupport := &SecondLineSupport{}
	slSupport.setNext(engineer)
	flSupport := &FirstLineSupport{}
	flSupport.setNext(slSupport)

	client := &Client{name: "Ivan"}
	flSupport.execute(client)
}

/*
	Цепочка обязанностей позволяет передавать запросы последовательно по цепочке обработчиков.
	Каждый следующий обработчик решает, может ли он обработать запрос сам либо передать дальше
	Избавляет от жёсткой привязки отправителя запроса к его получателю,
	позволяя выстраивать цепь из различных обработчиков динамически.
	Применимость:
	- когда нужно обрабатывать разообразные запросы несколькими способами, но заранее неизвестно,
	какие конкретно запросы будут приходить и какие обработчики понадобятся
	- когда важно, чтоб обработчики выполнялись строго в определенном порядке
	- когда набор объектов, способных обработать запрос должен задаваться динамически
	Плюсы:
	- уменьшает зависимоть между клиентом и обработчиком
	- реализует принцип единственной обязанности
	- реализует принцип открытости/закрытости
	Минусы:
	- запрос может остаться никем не обработанным
*/
