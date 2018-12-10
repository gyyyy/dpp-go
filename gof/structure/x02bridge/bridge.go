// Package x02bridge 桥接
package x02bridge

import (
	"log"
)

type Implementor interface {
	NewOperation()
}

type ConcreteImplementor1 struct {
}

func (i *ConcreteImplementor1) NewOperation() {
	log.Println("implementor[1] operate")
}

type ConcreteImplementor2 struct {
}

func (i *ConcreteImplementor2) NewOperation() {
	log.Println("implementor[2] operate")
}

type Abstraction interface {
	DoOperation()
}

type AbstractionImp struct {
	implementor Implementor
}

func (a *AbstractionImp) DoOperation() {
	log.Println("abstraction implement bridge implementor operate")
	a.implementor.NewOperation()
}

func NewAbstractionImp(implementor Implementor) *AbstractionImp {
	return &AbstractionImp{implementor}
}
