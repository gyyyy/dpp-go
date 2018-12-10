// Package x08nullobject 空对象
package x08nullobject

import (
	"log"
)

type AbstractOperation interface {
	Request()
}

type RealOperation struct {
	name string
}

func (o *RealOperation) Request() {
	log.Printf("real operation request: %s\n", o.name)
}

type NullOperation struct {
}

func (o *NullOperation) Request() {
	log.Println("null operation do nothing")
}

func NewOperation(name string) AbstractOperation {
	if name == "" {
		return new(NullOperation)
	}
	return &RealOperation{name}
}
