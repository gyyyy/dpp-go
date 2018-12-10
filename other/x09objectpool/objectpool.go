// Package x09objectpool 对象池
package x09objectpool

import (
	"log"
	"sync"
	"sync/atomic"
)

type Reusable interface {
	DoSomething()
}

type ConcreteReusable struct {
	ID int
}

func (r *ConcreteReusable) DoSomething() {
	log.Printf("reusable[%d] do something\n", r.ID)
}

type ReusablePool struct {
	counter uint32
	free    chan Reusable
}

func (p *ReusablePool) AcquireReusable() Reusable {
	select {
	case r := <-p.free:
		return r
	default:
		return &ConcreteReusable{int(atomic.AddUint32(&p.counter, 1))}
	}
}

func (p *ReusablePool) ReleaseReusable(reusable Reusable) {
	select {
	case p.free <- reusable:
	default:
	}
}

func (p *ReusablePool) Destroy() {
	close(p.free)
	for r := range p.free {
		log.Printf("release reusable[%d]\n", r.(*ConcreteReusable).ID)
	}
}

var (
	once sync.Once
	pool *ReusablePool
)

func Instance() *ReusablePool {
	once.Do(func() {
		pool = &ReusablePool{0, make(chan Reusable, 10)}
	})
	return pool
}
