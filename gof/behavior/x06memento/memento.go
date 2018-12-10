// Package x06memento 备忘录
package x06memento

type Memento struct {
	state string
}

func (m *Memento) State() string {
	return m.state
}

type Originator struct {
	state string
}

func (o *Originator) Update(state string) {
	o.state = state
}

func (o *Originator) Compare(state string) bool {
	return o.state == state
}

func (o *Originator) CreateMemento() *Memento {
	return &Memento{o.state}
}

func (o *Originator) SetMemento(memento *Memento) {
	o.state = memento.state
}

func NewOriginator(state string) *Originator {
	return &Originator{state}
}

type Caretaker struct {
	memento map[string]*Memento
}

func (c *Caretaker) Get(name string) (memento *Memento) {
	memento, _ = c.memento[name]
	return
}

func (c *Caretaker) Add(name string, memento Memento) {
	c.memento[name] = &memento
}

func NewCaretaker() *Caretaker {
	return &Caretaker{make(map[string]*Memento)}
}
