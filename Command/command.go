/*
 Command 命令模式：
    将一个请求封装为一个对象，
    从而使你可用不同的请求对客户进行参数化；
    对请求排队或者记录请求日志，以及支持可撤销的操作

 个人想法：Invoker维护请求队列（Command接口队列），通过一些函数可以添加、修改、执行请求队列，
         在每一种ConcreteCommand中有对该命令的执行体（Receiver），最终响应请求队列的命令
 作者：   HCLAC
 日期：   20170310
*/

package command

import (
	"fmt"
)

// 命令接口 -- 可以保存在请求队形中，方便请求队形处理命令，具体对命令的执行体在实现这个接口的类型结构体中保存着
type Command interface {
	Run()
}

// 请求队形，保存命令列表，在ExecuteCommand函数中遍历执行命令
type Invoker struct {
	comlist []Command
}

// 添加命令
func (i *Invoker) AddCommand(c Command) {
	if i == nil {
		return
	}
	i.comlist = append(i.comlist, c)
}

// 执行命令
func (i *Invoker) ExecuteCommand() {
	if i == nil {
		return
	}
	for _, val := range i.comlist {
		val.Run()
	}
}

func NewInvoker() *Invoker {
	return &Invoker{[]Command{}}
}

// 具体命令,实现Command接口，保存一个对该命令如何处理的执行体
type ConcreteCommandA struct {
	receiver ReceiverA
}

func (c *ConcreteCommandA) SetReceiver(r ReceiverA) {
	if c == nil {
		return
	}
	c.receiver = r
}

// 具体命令的执行体
func (c *ConcreteCommandA) Run() {
	if c == nil {
		return
	}
	c.receiver.Execute()
}

func NewConcreteCommandA() *ConcreteCommandA {
	return &ConcreteCommandA{}
}

// 针对ConcreteCommand，如何处理该命令
type ReceiverA struct {
}

func (r *ReceiverA) Execute() {
	if r == nil {
		return
	}
	fmt.Println("针对ConcreteCommandA，如何处理该命令")
}

func NewReceiverA() *ReceiverA {
	return &ReceiverA{}
}

//////////////////////////////////////////////////////////

// 具体命令,实现Command接口，保存一个对该命令如何处理的执行体
type ConcreteCommandB struct {
	receiver ReceiverB
}

func (c *ConcreteCommandB) SetReceiver(r ReceiverB) {
	if c == nil {
		return
	}
	c.receiver = r
}

// 具体命令的执行体
func (c *ConcreteCommandB) Run() {
	if c == nil {
		return
	}
	c.receiver.Execute()
}

func NewConcreteCommandB() *ConcreteCommandB {
	return &ConcreteCommandB{}
}

// 针对ConcreteCommandB，如何处理该命令
type ReceiverB struct {
}

func (r *ReceiverB) Execute() {
	if r == nil {
		return
	}
	fmt.Println("针对ConcreteCommandB，如何处理该命令")
}

func NewReceiverB() *ReceiverB {
	return &ReceiverB{}
}
