// Package x01chainofresponsibility 责任链
package x01chainofresponsibility

import (
	"log"
)

type Request struct {
	value int
}

func NewRequest(value int) *Request {
	return &Request{value}
}

type Handler interface {
	HandleRequest(*Request)
}

type defaultHandler struct {
	successor Handler
}

func (h *defaultHandler) SetSuccessor(successor Handler) {
	h.successor = successor
}

func (h *defaultHandler) HandleRequest(r *Request) {
	if h.successor != nil {
		h.successor.HandleRequest(r)
	}
}

type ConcreteHandler1 struct {
	defaultHandler
}

func (h *ConcreteHandler1) HandleRequest(r *Request) {
	if r.value < 10 {
		log.Println("handle by handler[1]")
		return
	}
	h.defaultHandler.HandleRequest(r)
}

type ConcreteHandler2 struct {
	defaultHandler
}

func (h *ConcreteHandler2) HandleRequest(r *Request) {
	if r.value < 100 {
		log.Println("handle by handler[2]")
		return
	}
	h.defaultHandler.HandleRequest(r)
}

type ConcreteHandler3 struct {
	defaultHandler
}

func (h *ConcreteHandler3) HandleRequest(r *Request) {
	if r.value < 1000 {
		log.Println("handle by handler[3]")
		return
	}
	h.defaultHandler.HandleRequest(r)
}
