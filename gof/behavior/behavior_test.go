package behavior

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/gyyyy/dpp/gof/behavior/x01chainofresponsibility"
	"github.com/gyyyy/dpp/gof/behavior/x02command"
	"github.com/gyyyy/dpp/gof/behavior/x03interpreter"
	"github.com/gyyyy/dpp/gof/behavior/x04iterator"
	"github.com/gyyyy/dpp/gof/behavior/x05mediator"
	"github.com/gyyyy/dpp/gof/behavior/x06memento"
	"github.com/gyyyy/dpp/gof/behavior/x07observer"
	"github.com/gyyyy/dpp/gof/behavior/x08state"
	"github.com/gyyyy/dpp/gof/behavior/x09strategy"
	"github.com/gyyyy/dpp/gof/behavior/x10templatemethod"
	"github.com/gyyyy/dpp/gof/behavior/x11visitor"
)

func TestChainOfResponsibility(t *testing.T) {
	handler1 := new(x01chainofresponsibility.ConcreteHandler1)
	handler2 := new(x01chainofresponsibility.ConcreteHandler2)
	handler3 := new(x01chainofresponsibility.ConcreteHandler3)
	handler1.SetSuccessor(handler2)
	handler2.SetSuccessor(handler3)
	handler1.HandleRequest(x01chainofresponsibility.NewRequest(1))
	handler1.HandleRequest(x01chainofresponsibility.NewRequest(11))
	handler1.HandleRequest(x01chainofresponsibility.NewRequest(111))
}

func TestCommand(t *testing.T) {
	receiver := new(x02command.Receiver)
	invoker := x02command.NewInvoker()
	go func() {
		defer invoker.Over()
		for i := 0; i < 20; i++ {
			cmd := x02command.NewConcreteCommand(fmt.Sprintf("test%d", i+1))
			cmd.SetReceiver(receiver)
			invoker.SetCommand(cmd)
		}
	}()
	invoker.Invoke()
}

func TestInterpreter(t *testing.T) {
	ctx01 := x03interpreter.NewContext("大哥")
	new(x03interpreter.TerminalExpression).Interpret(ctx01)
	if ctx01.Output != url.QueryEscape(ctx01.Input)+"%0A" {
		t.Fail()
	}
	ctx02 := x03interpreter.NewContext("大哥123\n")
	new(x03interpreter.NonTerminalExpression).Interpret(ctx02)
	if ctx02.Output+"%0A" != url.QueryEscape(ctx02.Input) {
		t.Fail()
	}
}

func TestIterator(t *testing.T) {
	m := make([]interface{}, 0)
	for i := 0x20; i < 0x07f; i++ {
		m = append(m, string(i))
	}
	iterator := x04iterator.NewConcreteAggregate(m).Iterator()
	for i := 0x20; iterator.HasNext(); i++ {
		if iterator.Next() != string(i) {
			t.Fail()
		}
	}
}

func TestMediator(t *testing.T) {
	colleague1 := x05mediator.NewConcreteColleague("test1")
	colleague2 := x05mediator.NewConcreteColleague("test2")
	mediator := x05mediator.NewConcreteMediator(colleague1, colleague2)
	colleague1.Send(x05mediator.NewMessage("test2", "test1to2"), mediator)
	colleague2.Send(x05mediator.NewMessage("test2", "test2to2"), mediator)
	colleague2.Send(x05mediator.NewMessage("test3", "你好"), mediator)
}

func TestMemento(t *testing.T) {
	var (
		name, old, new = "test", "state1", "state2"
		originator     = x06memento.NewOriginator(old)
		caretaker      = x06memento.NewCaretaker()
	)
	caretaker.Add(name, *originator.CreateMemento())
	originator.Update(new)
	originator.SetMemento(caretaker.Get(name))
	if !originator.Compare(old) {
		t.Fail()
	}
}

func TestObserver(t *testing.T) {
	observable := x07observer.NewConcreteObservable()
	observable.Attach(x07observer.NewConcreteObserver("test1"))
	observable.Attach(x07observer.NewConcreteObserver("test2"))
	observable.DoSomething()
}

func TestState(t *testing.T) {
	ctx := new(x08state.Context)
	ctx.SetState(new(x08state.ConcreteStateB))
	ctx.Request()
}

func TestStrategy(t *testing.T) {
	ctx := x09strategy.NewContext(new(x09strategy.ConcreteStrategyA))
	ctx.DoSomething()
	ctx.ChangeStrategy(new(x09strategy.ConcreteStrategyB))
	ctx.DoSomething()
}

func TestTemplateMethod(t *testing.T) {
	x10templatemethod.NewTemplateClass(new(x10templatemethod.ConcreteClass)).TemplateMethod()
}

func TestVisitor(t *testing.T) {
	object := x11visitor.NewObjectStructure(
		new(x11visitor.ConcreteVisitableA),
		new(x11visitor.ConcreteVisitableB),
	)
	object.Accept(new(x11visitor.ConcreteVisitorA))
}
