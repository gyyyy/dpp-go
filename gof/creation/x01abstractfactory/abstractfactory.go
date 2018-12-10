// Package x01abstractfactory 抽象工厂
package x01abstractfactory

import (
	"log"
)

type AbstractProductA interface {
	DoSomethingByA()
}

type ProductA1 struct {
}

func (p *ProductA1) DoSomethingByA() {
	log.Println("product[A1] do something")
}

type ProductA2 struct {
}

func (p *ProductA2) DoSomethingByA() {
	log.Println("product[A2] do something")
}

type AbstractProductB interface {
	DoSomethingByB()
}

type ProductB1 struct {
}

func (p *ProductB1) DoSomethingByB() {
	log.Println("product[B1] do something")
}

type ProductB2 struct {
}

func (p *ProductB2) DoSomethingByB() {
	log.Println("product[B2] do something")
}

type AbstractFactory interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}

type ConcreteFactory1 struct {
}

func (f *ConcreteFactory1) CreateProductA() AbstractProductA {
	return new(ProductA1)
}

func (f *ConcreteFactory1) CreateProductB() AbstractProductB {
	return new(ProductB1)
}

type ConcreteFactory2 struct {
}

func (f *ConcreteFactory2) CreateProductA() AbstractProductA {
	return new(ProductA2)
}

func (f *ConcreteFactory2) CreateProductB() AbstractProductB {
	return new(ProductB2)
}

func Factory(name string) AbstractFactory {
	switch name {
	case "a":
		return new(ConcreteFactory1)
	case "b":
		return new(ConcreteFactory2)
	default:
		return nil
	}
}
