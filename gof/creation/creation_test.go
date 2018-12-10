package creation

import (
	"testing"

	"github.com/gyyyy/dpp/gof/creation/x01abstractfactory"
	"github.com/gyyyy/dpp/gof/creation/x02builder"
	"github.com/gyyyy/dpp/gof/creation/x03factorymethod"
	"github.com/gyyyy/dpp/gof/creation/x04prototype"
	"github.com/gyyyy/dpp/gof/creation/x05singleton"
)

func TestAbstractFactory(t *testing.T) {
	x01abstractfactory.Factory("a").CreateProductA().DoSomethingByA()
}

func TestBuilder(t *testing.T) {
	builder := new(x02builder.ConcreteBuilder)
	x02builder.NewDirector(builder).Construct()
	if builder.Result().N() != 1024 {
		t.Fail()
	}
}

func TestFactoryMethod(t *testing.T) {
	new(x03factorymethod.ConcreteFactory).DoSomething()
}

func TestPrototype(t *testing.T) {
	prototype1 := x04prototype.NewConcretePrototype("test1")
	prototype2 := prototype1.Clone().(*x04prototype.ConcretePrototype)
	if prototype2.Name = "test2"; prototype1.Name == prototype2.Name {
		t.Fail()
	}
}

func TestSingleton(t *testing.T) {
	x05singleton.Instance().DoSomething()
}
