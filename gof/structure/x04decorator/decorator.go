// Package x04decorator 装饰器
package x04decorator

import (
	"log"
)

type Component interface {
	DoOperation()
}

type ConcreteComponent struct {
}

func (c *ConcreteComponent) DoOperation() {
	log.Println("component operate")
}

type Decorator struct {
	component Component
}

func (d *Decorator) DoOperation() {
	d.component.DoOperation()
}

type ConcreteDecoratorExtendingFunctionality struct {
	Decorator
}

func (d *ConcreteDecoratorExtendingFunctionality) DoAdditionalOperation() {
	log.Println("decorator[func] additional operate")
}

func (d *ConcreteDecoratorExtendingFunctionality) DoOperation() {
	d.DoAdditionalOperation()
	d.component.DoOperation()
}

func NewConcreteDecoratorExtendingFunctionality(component Component) *ConcreteDecoratorExtendingFunctionality {
	return &ConcreteDecoratorExtendingFunctionality{Decorator{component}}
}

type ConcreteDecoratorExtendingState struct {
	Decorator
	state string
}

func (d *ConcreteDecoratorExtendingState) State() string {
	return d.state
}

func (d *ConcreteDecoratorExtendingState) DoOperation() {
	d.state = "new state"
	log.Println("decorator[state] change state")
	d.component.DoOperation()
}

func NewConcreteDecoratorExtendingState(component Component, state string) *ConcreteDecoratorExtendingState {
	return &ConcreteDecoratorExtendingState{Decorator{component}, state}
}
