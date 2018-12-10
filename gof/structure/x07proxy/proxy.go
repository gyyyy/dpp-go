// Package x07proxy 代理
package x07proxy

import (
	"log"
)

type Subject interface {
	DoOperation()
}

type RealSubject struct {
}

func (s *RealSubject) DoOperation() {
	log.Println("real subject operate")
}

type Proxy struct {
	subject Subject
}

func (p *Proxy) DoOperation() {
	log.Println("proxy start")
	p.subject.DoOperation()
	log.Println("proxy end")
}

func NewProxy(subject Subject) *Proxy {
	return &Proxy{subject}
}
