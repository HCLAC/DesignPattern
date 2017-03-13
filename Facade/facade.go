/*
  Facade 外观模式：
        为子系统中的一组接口提供一个一致的界面，此模式定义了一个高层接口，
		这个接口使得这一子系统更加容易使用（投资：基金，股票，房产）
 个人想法：中介者模式、外观模式：每个对象都保存一份中介者对象，
        在和其他对象交互时，通过中介者来完成，外观模式：外观中保存了一堆对象，
		这些对象或者是组成某个子系统的，将其封装在外观对象中，给客户端一种只有一个对象的感觉，
		一个是结构型模式，一个是行为性模式
 作者：   HCLAC
 日期：   20170309
*/
package facade

import (
	"fmt"
)

type facade struct {
	one   funcone
	two   functwo
	three functhree
}

func (f facade) OutOne() {
	f.one.Out()
	f.three.Out()
}

func (f facade) OutTwo() {
	f.two.Out()
	f.three.Out()
}

type funcone struct {
	str string
}

func (f funcone) Out() {
	fmt.Println("funcone", f.str)
}

type functwo struct {
	i int
}

func (f functwo) Out() {
	fmt.Println("functwo", f.i)
}

type functhree struct {
	f float32
}

func (f functhree) Out() {
	fmt.Println("functhree", f.f)
}

func NewFacade(i int, f float32, str string) *facade {
	return &facade{funcone{str}, functwo{i}, functhree{f}}
}
