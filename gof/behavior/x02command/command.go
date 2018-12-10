// Package x02command 命令
package x02command

import (
	"log"
)

type Command interface {
	Execute()
}

type ConcreteCommand struct {
	receiver *Receiver
	s        string
}

func (c *ConcreteCommand) SetReceiver(receiver *Receiver) {
	c.receiver = receiver
}

func (c *ConcreteCommand) Execute() {
	c.receiver.Action(c.s)
}

func NewConcreteCommand(s string) *ConcreteCommand {
	return &ConcreteCommand{s: s}
}

type Invoker struct {
	queue chan Command
}

func (i *Invoker) InvokeNow(cmd Command) {
	cmd.Execute()
}

func (i *Invoker) SetCommand(cmd Command) {
	i.queue <- cmd
}

func (i *Invoker) Invoke() {
	for cmd := range i.queue {
		cmd.Execute()
	}
}

func (i *Invoker) Over() {
	close(i.queue)
}

func NewInvoker() *Invoker {
	return &Invoker{make(chan Command, 10)}
}

type Receiver struct {
}

func (r *Receiver) Action(s string) {
	log.Printf("receiver receive the command: %s\n", s)
}
