// Package x11visitor 访问者
package x11visitor

import (
	"log"
)

type Visitable interface {
	Accept(Visitor)
}

type ConcreteVisitableA struct {
}

func (v *ConcreteVisitableA) Accept(visitor Visitor) {
	visitor.Visit(v)
}

func (v *ConcreteVisitableA) DoSomething() {
	log.Println("visitable[A] do something")
}

type ConcreteVisitableB struct {
}

func (v *ConcreteVisitableB) Accept(visitor Visitor) {
	visitor.Visit(v)
}

func (v *ConcreteVisitableB) DoSomething() {
	log.Println("visitable[A] do something")
}

type ObjectStructure struct {
	visitable []Visitable
}

func (o *ObjectStructure) Accept(visitor Visitor) {
	for _, v := range o.visitable {
		v.Accept(visitor)
	}
}

func NewObjectStructure(visitable ...Visitable) *ObjectStructure {
	return &ObjectStructure{visitable}
}

type Visitor interface {
	Visit(Visitable)
}

type ConcreteVisitorA struct {
}

func (v *ConcreteVisitorA) Visit(visitable Visitable) {
	switch v := visitable.(type) {
	case (*ConcreteVisitableA):
		log.Println("visitor[A] visit visitable[A]")
		v.DoSomething()
	case (*ConcreteVisitableB):
		log.Println("visitor[A] visit visitable[B]")
		v.DoSomething()
	}
}

type ConcreteVisitorB struct {
}

func (v *ConcreteVisitorB) Visit(visitable Visitable) {
	switch v := visitable.(type) {
	case (*ConcreteVisitableA):
		log.Println("visitor[B] visit visitable[A]")
		v.DoSomething()
	case (*ConcreteVisitableB):
		log.Println("visitor[B] visit visitable[B]")
		v.DoSomething()
	}
}
