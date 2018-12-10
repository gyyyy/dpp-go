// Package x09strategy 策略
package x09strategy

import (
	"log"
)

type Strategy interface {
	BehaviorInterface()
}

type ConcreteStrategyA struct {
}

func (s *ConcreteStrategyA) BehaviorInterface() {
	log.Println("use strategy[A]")
}

type ConcreteStrategyB struct {
}

func (s *ConcreteStrategyB) BehaviorInterface() {
	log.Println("use strategy[B]")
}

type Context struct {
	strategy Strategy
}

func (c *Context) ChangeStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) DoSomething() {
	c.strategy.BehaviorInterface()
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy}
}
