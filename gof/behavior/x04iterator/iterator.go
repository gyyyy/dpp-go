// Package x04iterator 迭代器
package x04iterator

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type ConcreteIterator struct {
	queue chan interface{}
}

func (i *ConcreteIterator) HasNext() bool {
	return len(i.queue) != 0
}

func (i *ConcreteIterator) Next() interface{} {
	select {
	case v := <-i.queue:
		return v
	default:
		return nil
	}
}

type Aggregate interface {
	Iterator() Iterator
}

type ConcreteAggregate struct {
	item []interface{}
}

func (a *ConcreteAggregate) Iterator() Iterator {
	item := make(chan interface{}, len(a.item))
	for _, v := range a.item {
		item <- v
	}
	close(item)
	return &ConcreteIterator{item}
}

func NewConcreteAggregate(item []interface{}) *ConcreteAggregate {
	return &ConcreteAggregate{item}
}
