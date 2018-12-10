// Package x06flyweight 享元
package x06flyweight

import (
	"log"
	"sync"
)

type Flyweight interface {
	DoOperation(string)
}

type ConcreteFlyweight struct {
	intrinsicState string
}

func (f *ConcreteFlyweight) DoOperation(extrinsicState string) {
	log.Printf("flyweight operate with extrinsic state: %s\n", extrinsicState)
	log.Printf("flyweight hold intrinsic state: %s\n", f.intrinsicState)
}

var (
	once      sync.Once
	flyweight Flyweight
)

func Instance() Flyweight {
	once.Do(func() {
		flyweight = &ConcreteFlyweight{"intrinsic state"}
	})
	return flyweight
}
