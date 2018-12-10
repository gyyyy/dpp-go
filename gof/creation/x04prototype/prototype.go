// Package x04prototype 原型
package x04prototype

type Prototype interface {
	Clone() interface{}
}

type ConcretePrototype struct {
	Name string
}

func (p *ConcretePrototype) Clone() interface{} {
	tmp := *p
	return &tmp
}

func NewConcretePrototype(name string) *ConcretePrototype {
	return &ConcretePrototype{name}
}
