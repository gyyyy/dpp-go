// Package x03composite 组合
package x03composite

import (
	"log"
)

type Component interface {
	DoOperation()
}

type Leaf struct {
	name string
}

func (l *Leaf) DoOperation() {
	log.Printf("leaf[%s] operate\n", l.name)
}

func NewLeaf(name string) *Leaf {
	return &Leaf{name}
}

type Composite struct {
	name      string
	component []Component
}

func (c *Composite) AddComponent(component Component) {
	c.component = append(c.component, component)
}

func (c *Composite) RemoveComponent(i int) {
	if i < 0 || i >= len(c.component) {
		return
	}
	c.component = append(c.component[:i], c.component[i+1:]...)
}

func (c *Composite) GetChild(i int) Component {
	if i < 0 || i >= len(c.component) {
		return nil
	}
	return c.component[i]
}

func (c *Composite) DoOperation() {
	log.Printf("composite[%s] operate\n", c.name)
	for _, v := range c.component {
		v.DoOperation()
	}
}

func NewComposite(name string, component ...Component) *Composite {
	return &Composite{name, component}
}
