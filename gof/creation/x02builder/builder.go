// Package x02builder 生成器
package x02builder

type Product struct {
	n int
}

func (p *Product) N() int {
	return p.n
}

type Builder interface {
	BuildPart()
}

type ConcreteBuilder struct {
	n int
}

func (b *ConcreteBuilder) BuildPart() {
	b.n = 1024
}

func (b *ConcreteBuilder) Result() *Product {
	return &Product{b.n}
}

type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.BuildPart()
}

func NewDirector(builder Builder) *Director {
	return &Director{builder}
}
