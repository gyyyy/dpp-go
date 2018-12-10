// Package x01adapter 适配器
package x01adapter

import (
	"log"
)

type Target interface {
	Request()
}

type Adaptee interface {
	SpecificRequest()
}

type ConcreteAdaptee struct {
}

func (a *ConcreteAdaptee) SpecificRequest() {
	log.Println("adaptee specific request")
}

type Adapter struct {
	adaptee Adaptee
}

func (a *Adapter) Request() {
	log.Println("adapter adapte adaptee specific request to request")
	a.adaptee.SpecificRequest()
}

func NewAdapter(adaptee Adaptee) *Adapter {
	return &Adapter{adaptee}
}
