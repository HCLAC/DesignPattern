/*
  Decorator 装饰模式：
        动态地给一个对象添加一些额外的职责，就增加功能来说，装饰模式比生成子类更为灵活。
 个人想法：注意装饰器内部维护一个对象，所有装饰器的子类在操作时，必须调用父类的函数，一直从下到上再到下的感觉
 作者：   HCLAC
 日期：   20170308
*/
package decorator

import (
	"fmt"
)

type person struct {
	Name string
}

func (p *person) show() {
	if p == nil {
		return
	}
	fmt.Println("姓名：", p.Name)
}

type AbsstractPerson interface {
	show()
}
type Decorator struct {
	AbsstractPerson
}

func (d *Decorator) SetDecorator(component AbsstractPerson) {
	if d == nil {
		return
	}
	d.AbsstractPerson = component
}

func (d *Decorator) show() {
	if d == nil {
		return
	}
	if d.AbsstractPerson != nil {
		d.AbsstractPerson.show()
	}
}

type TShirts struct {
	Decorator
}

func (t *TShirts) show() {
	if t == nil {
		return
	}
	t.Decorator.show()
	fmt.Println("T恤")
}

type BigTrouser struct {
	Decorator
}

func (b *BigTrouser) show() {
	if b == nil {
		return
	}
	b.Decorator.show()
	fmt.Println("大裤衩")
}

type Sneakers struct {
	Decorator
}

func (b *Sneakers) show() {
	if b == nil {
		return
	}
	b.Decorator.show()
	fmt.Println("破球鞋")
}
