// Package x05facade 外观
package x05facade

import (
	"log"
)

type Component1 struct {
}

func (c *Component1) DoSomething() {
	log.Println("component[1] do something")
}

type Component2 struct {
}

func (c *Component2) DoSomething() {
	log.Println("component[2] do something")
}

type Component3 struct {
}

func (c *Component3) DoSomething() {
	log.Println("component[3] do something")
}

type Facade struct {
	component1 *Component1
	component2 *Component2
	component3 *Component3
}

func (f *Facade) DoSomething() {
	log.Println("facade do something")
	f.component1.DoSomething()
	f.component2.DoSomething()
	f.component3.DoSomething()
}

func NewFacade(component1 *Component1, component2 *Component2, component3 *Component3) *Facade {
	return &Facade{component1, component2, component3}
}
