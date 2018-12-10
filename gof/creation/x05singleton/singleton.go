// Package x05singleton 单例
package x05singleton

import (
	"log"
	"sync"
)

type Singleton struct {
}

func (s *Singleton) DoSomething() {
	log.Println("singleton do something")
}

var (
	once     sync.Once
	instance *Singleton
)

func Instance() *Singleton {
	once.Do(func() {
		instance = new(Singleton)
	})
	return instance
}
