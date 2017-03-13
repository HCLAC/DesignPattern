/*
 Visitor 访问者模式：
        表示一个作用于某对象结构中的各元素的操作，
		它使你可以在不改变各元素的类的前提下定义作用于这些元素的新操作
 个人想法：
 作者：   HCLAC
 日期：   20170309
*/

package visitor

import (
	"fmt"
)

// 访问接口
type IVisitor interface {
	VisitConcreteElementA(ConcreteElementA)
	VisitConcreteElementB(ConcreteElementB)
}

// 具体访问者A
type ConcreteVisitorA struct {
	name string
}

func (c *ConcreteVisitorA) VisitConcreteElementA(ce ConcreteElementA) {
	if c == nil {
		return
	}
	fmt.Println(ce.name, c.name)
	ce.OperatorA()
}

func (c *ConcreteVisitorA) VisitConcreteElementB(ce ConcreteElementB) {
	if c == nil {
		return
	}
	fmt.Println(ce.name, c.name)
	ce.OperatorB()
}

// 具体访问者B
type ConcreteVisitorB struct {
	name string
}

func (c *ConcreteVisitorB) VisitConcreteElementA(ce ConcreteElementA) {
	if c == nil {
		return
	}
	fmt.Println(ce.name, c.name)
	ce.OperatorA()
}

func (c *ConcreteVisitorB) VisitConcreteElementB(ce ConcreteElementB) {
	if c == nil {
		return
	}
	fmt.Println(ce.name, c.name)
	ce.OperatorB()
}

// 元素接口
type IElement interface {
	Accept(IVisitor)
}

// 具体元素A
type ConcreteElementA struct {
	name string
}

func (c *ConcreteElementA) Accept(visitor IVisitor) {
	if c == nil {
		return
	}
	visitor.VisitConcreteElementA(*c)
}
func (c *ConcreteElementA) OperatorA() {
	if c == nil {
		return
	}
	fmt.Println("OperatorA")
}

// 具体元素B
type ConcreteElementB struct {
	name string
}

func (c *ConcreteElementB) Accept(visitor IVisitor) {
	if c == nil {
		return
	}
	visitor.VisitConcreteElementB(*c)
}
func (c *ConcreteElementB) OperatorB() {
	if c == nil {
		return
	}
	fmt.Println("OperatorB")
}

// 维护元素集合
type ObjectStructure struct {
	list []IElement
}

func (o *ObjectStructure) Attach(e IElement) {
	if o == nil || e == nil {
		return
	}
	o.list = append(o.list, e)
}

func (o *ObjectStructure) Detach(e IElement) {
	if o == nil || e == nil {
		return
	}
	for i, val := range o.list {
		if val == e {
			o.list = append(o.list[:i], o.list[i+1:]...)
			break
		}
	}
}

func (o *ObjectStructure) Accept(v IVisitor) {
	if o == nil {
		return
	}
	for _, val := range o.list {
		val.Accept(v)
	}
}
