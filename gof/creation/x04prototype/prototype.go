// Package x04prototype 原型
package x04prototype

type Prototype interface {
	Clone() *ConcretePrototype
}

type ConcretePrototype struct {
	Name string
}

func (p *ConcretePrototype) Clone() *ConcretePrototype {
	tmp := *p
	return &tmp
}

func NewConcretePrototype(name string) *ConcretePrototype {
	return &ConcretePrototype{name}
}
