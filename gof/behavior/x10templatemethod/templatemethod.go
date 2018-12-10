// Package x10templatemethod 模板方法
package x10templatemethod

import (
	"log"
)

type AbstractClass interface {
	PrimitiveOperationA()
	PrimitiveOperationB()
}

type ConcreteClass struct {
	AbstractClass
}

func (c *ConcreteClass) PrimitiveOperationA() {
	log.Println("class primitive operate A")
}

func (c *ConcreteClass) PrimitiveOperationB() {
	log.Println("class primitive operate B")
}

type TemplateClass struct {
	class AbstractClass
}

func (c *TemplateClass) TemplateMethod() {
	log.Println("template method start")
	c.class.PrimitiveOperationA()
	c.class.PrimitiveOperationB()
	log.Println("template method end")
}

func NewTemplateClass(class AbstractClass) *TemplateClass {
	return &TemplateClass{class}
}
