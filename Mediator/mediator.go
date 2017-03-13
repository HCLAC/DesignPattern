/*
 Mediator 中介者模式：
        用一个中介对象来封装一系列的对象交互。
        中介这使各对象不需要显式地相互引用，从而使其耦合松散，
        而且可以独立地改变它们之间的交互。
 个人想法：每个对象都有一个中介者对象，发生变化时，通知中介者，由中介者判断通知其他的对象。
 作者：   HCLAC
*/

package Mediator

import (
	"fmt"
)

// 中介者接口
type IMediator interface {
	Send(string, IColleague)
}

// 实现中介者接口的基本类型
type Mediator struct {
}

// 具体的中介者
type ConcreteMediator struct {
	Mediator
	colleagues []IColleague
}

func (m *ConcreteMediator) AddColleague(c IColleague) {
	if m == nil {
		return
	}
	m.colleagues = append(m.colleagues, c)
}

func (m *ConcreteMediator) Send(message string, c IColleague) {
	if m == nil {
		return
	}
	for _, val := range m.colleagues {
		if c == val {
			continue
		}
		val.Notify(message)
	}
}

func NewConcreteMediator() *ConcreteMediator {
	return &ConcreteMediator{}
}

// 合作者接口
type IColleague interface {
	Send(string)
	Notify(string)
}

// 实现合作者接口的基本类型
type Colleague struct {
	mediator IMediator
}

// 具体合作者对象A
type ConcreteColleageA struct {
	Colleague
}

func (c *ConcreteColleageA) Notify(message string) {
	if c == nil {
		return
	}
	fmt.Println("ConcreteColleageA get message:", message)
}

func (c *ConcreteColleageA) Send(message string) {
	if c == nil {
		return
	}
	c.mediator.Send(message, c)

}
func NewConcreteColleageA(mediator IMediator) *ConcreteColleageA {
	return &ConcreteColleageA{Colleague{mediator}}
}

// 具体合作者对象B
type ConcreteColleageB struct {
	Colleague
}

func (c *ConcreteColleageB) Notify(message string) {
	if c == nil {
		return
	}
	fmt.Println("ConcreteColleageB get message:", message)
}
func (c *ConcreteColleageB) Send(message string) {
	if c == nil {
		return
	}
	c.mediator.Send(message, c)

}
func NewConcreteColleageB(mediator IMediator) *ConcreteColleageB {
	return &ConcreteColleageB{Colleague{mediator}}
}
