// Package x08state 状态
package x08state

import (
	"log"
)

type State interface {
	Handle()
}

type ConcreteStateA struct {
}

func (s *ConcreteStateA) Handle() {
	log.Println("state[A] handle")
}

type ConcreteStateB struct {
}

func (s *ConcreteStateB) Handle() {
	log.Println("state[B] handle")
}

type ConcreteStateC struct {
}

func (s *ConcreteStateC) Handle() {
	log.Println("state[C] handle")
}

type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle()
	for i := 0; i < 5; i++ {
		switch i % 3 {
		case 0:
			c.state = new(ConcreteStateA)
		case 1:
			c.state = new(ConcreteStateB)
		case 2:
			c.state = new(ConcreteStateC)
		}
		c.state.Handle()
	}
}
