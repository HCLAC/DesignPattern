/*
 Chain Of Responsibility 职责链模式：
        使多个对象都有机会处理请求，从而避免请求的发送者和接收者之间的耦合关系。
        将这些对象连成一条链，并沿着这条链传递该请求，直到有一个对象处理它为止

 个人想法：MFC的消息队列机制
 作者：   HCLAC
 日期：   20170309
*/
package chainofresponsibility

import (
	"fmt"
)

const (
	constHandler = iota
	constHandlerA
	constHandlerB
)

// 处理请求接口
type IHandler interface {
	SetSuccessor(IHandler)
	HandleRequest(int) int
}

// 实现处理请求的接口的基本结构体类型
type Handler struct {
	successor IHandler // 继承者
}

func (h *Handler) SetSuccessor(i IHandler) {
	if h == nil {
		return
	}
	h.successor = i
}

// 具体处理结构体，这里简单处理int类型的请求，判断是否在[1-10]之间，是：处理，否：交给successor处理
type ConcreteHandlerA struct {
	Handler
}

func (c *ConcreteHandlerA) HandleRequest(req int) int {
	if c == nil {
		return constHandler
	}
	if req > 0 && req < 11 {
		fmt.Println("ConcreteHandlerA可以处理这个请求")
		return constHandlerA
	} else if c.successor != nil {
		return c.successor.HandleRequest(req)
	}
	return constHandler
}

func NewConcreteHandlerA() *ConcreteHandlerA {
	return &ConcreteHandlerA{}
}

// 具体处理结构体，这里简单处理int类型的请求，判断是否在[11-20]之间，是：处理，否：交给successor处理
type ConcreteHandlerB struct {
	Handler
}

func (c *ConcreteHandlerB) HandleRequest(req int) int {
	if c == nil {
		return constHandler
	}
	if req > 10 && req < 21 {
		fmt.Println("ConcreteHandlerB可以处理这个请求")
		return constHandlerB
	} else if c.successor != nil {
		return c.successor.HandleRequest(req)
	}
	return constHandler
}

func NewConcreteHandlerB() *ConcreteHandlerB {
	return &ConcreteHandlerB{}
}
