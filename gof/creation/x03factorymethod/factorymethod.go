// Package x03factorymethod 工厂方法
package x03factorymethod

import (
	"log"
)

type Product interface {
	DoSomething()
}

type ConcreteProduct struct {
}

func (p *ConcreteProduct) DoSomething() {
	log.Println("product do something")
}

type Factory interface {
	FactoryMethod() Product
}

type ConcreteFactory struct {
}

func (f *ConcreteFactory) FactoryMethod() Product {
	return new(ConcreteProduct)
}

func (f *ConcreteFactory) DoSomething() {
	f.FactoryMethod().DoSomething()
}
