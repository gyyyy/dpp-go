// Package x07observer 观察者
package x07observer

import (
	"log"
)

type Observer interface {
	Name() string
	Update()
}

type ConcreteObserver struct {
	name string
}

func (o *ConcreteObserver) Name() string {
	return o.name
}

func (o *ConcreteObserver) Update() {
	log.Printf("%s update\n", o.name)
}

func NewConcreteObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{name}
}

type Observable interface {
	Attach(Observer)
	Detach(Observer)
	Notify()
}

type ConcreteObservable struct {
	observer map[string]Observer
}

func (o *ConcreteObservable) Attach(observer Observer) {
	o.observer[observer.Name()] = observer
}

func (o *ConcreteObservable) Detach(observer Observer) {
	delete(o.observer, observer.Name())
}

func (o *ConcreteObservable) Notify() {
	for _, v := range o.observer {
		v.Update()
	}
}

func (o *ConcreteObservable) DoSomething() {
	log.Println("observable do something")
	o.Notify()
}

func NewConcreteObservable() *ConcreteObservable {
	return &ConcreteObservable{make(map[string]Observer)}
}
