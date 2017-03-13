package Mediator

import (
	"testing"
)

func TestMediator(t *testing.T) {
	m := NewConcreteMediator()
	ca := NewConcreteColleageA(m)
	cb := NewConcreteColleageB(m)

	m.AddColleague(ca)
	m.AddColleague(cb)

	ca.Send("ca")
	cb.Send("cb")
}
