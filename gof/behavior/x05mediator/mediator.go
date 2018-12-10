// Package x05mediator 中介者
package x05mediator

import (
	"log"
)

type Message struct {
	From Colleague
	To   string
	Data string
}

func NewMessage(to, data string) *Message {
	return &Message{To: to, Data: data}
}

type Colleague interface {
	Name() string
	Send(*Message, Mediator)
	Receive(*Message)
}

type ConcreteColleague struct {
	name string
}

func (c *ConcreteColleague) Name() string {
	return c.name
}

func (c *ConcreteColleague) Send(msg *Message, mediator Mediator) {
	msg.From = c
	mediator.Mediate(msg)
}

func (c *ConcreteColleague) Receive(msg *Message) {
	log.Printf("from %s to %s: %s\n", msg.From.Name(), msg.To, msg.Data)
}

func NewConcreteColleague(name string) *ConcreteColleague {
	return &ConcreteColleague{name}
}

type Mediator interface {
	Mediate(*Message)
}

type ConcreteMediator struct {
	colleague map[string]Colleague
}

func (m *ConcreteMediator) Mediate(msg *Message) {
	if msg.To == msg.From.Name() {
		log.Printf("can not send message to yourself[%s]: %s\n", msg.To, msg.Data)
		return
	}
	to, ok := m.colleague[msg.To]
	if !ok {
		log.Printf("not found colleague: %s\n", msg.To)
		return
	}
	to.Receive(msg)
}

func NewConcreteMediator(colleague ...Colleague) *ConcreteMediator {
	m := make(map[string]Colleague)
	for _, c := range colleague {
		m[c.Name()] = c
	}
	return &ConcreteMediator{m}
}
