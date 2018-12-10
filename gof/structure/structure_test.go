package structure

import (
	"testing"

	"github.com/gyyyy/dpp/gof/structure/x01adapter"
	"github.com/gyyyy/dpp/gof/structure/x02bridge"
	"github.com/gyyyy/dpp/gof/structure/x03composite"
	"github.com/gyyyy/dpp/gof/structure/x04decorator"
	"github.com/gyyyy/dpp/gof/structure/x05facade"
	"github.com/gyyyy/dpp/gof/structure/x06flyweight"
	"github.com/gyyyy/dpp/gof/structure/x07proxy"
)

func TestAdapter(t *testing.T) {
	adaptee := new(x01adapter.ConcreteAdaptee)
	adaptee.SpecificRequest()
	x01adapter.NewAdapter(adaptee).Request()
}

func TestBridge(t *testing.T) {
	implementor := new(x02bridge.ConcreteImplementor1)
	implementor.NewOperation()
	x02bridge.NewAbstractionImp(implementor).DoOperation()
}

func TestComposite(t *testing.T) {
	tree := []x03composite.Component{
		x03composite.NewLeaf("l1"),
		x03composite.NewComposite("c1",
			x03composite.NewComposite("c2",
				x03composite.NewLeaf("l2"),
				x03composite.NewLeaf("l3"),
			),
			x03composite.NewLeaf("l4")),
	}
	for _, v := range tree {
		v.DoOperation()
	}
}

func TestDecorator(t *testing.T) {
	component := new(x04decorator.ConcreteComponent)
	component.DoOperation()
	decorator1 := x04decorator.NewConcreteDecoratorExtendingFunctionality(component)
	decorator1.DoOperation()
	decorator2 := x04decorator.NewConcreteDecoratorExtendingState(component, "test")
	old := decorator2.State()
	decorator2.DoOperation()
	if decorator2.State() == old {
		t.Fail()
	}
}

func TestFacade(t *testing.T) {
	facade := x05facade.NewFacade(
		new(x05facade.Component1),
		new(x05facade.Component2),
		new(x05facade.Component3),
	)
	facade.DoSomething()
}

func TestFlyweight(t *testing.T) {
	flyweight := x06flyweight.Instance()
	flyweight.DoOperation("test1")
	flyweight.DoOperation("test2")
}

func TestProxy(t *testing.T) {
	real := new(x07proxy.RealSubject)
	real.DoOperation()
	proxy := x07proxy.NewProxy(real)
	proxy.DoOperation()
}
