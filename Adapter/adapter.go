/*
 Adapter 适配器模式：
        将一个类的接口转换成客户端希望的另一个接口。
		适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作
 个人想法：代理和适配器：代理和代理的对象接口一致，客户端不知道代理对象，
        而适配器是客户端想要适配器的接口，适配器对象的接口和客户端想要的不一样，
		适配器将适配器对象的接口封装一下，改成客户端想要的接口
 作者：   HCLAC
 日期：   20170306
*/

package adapter

import (
	"fmt"
)

type Player interface {
	attack()
	defense()
}

type Forwards struct {
	name string
}

func (f *Forwards) attack() {
	if f == nil {
		return
	}
	fmt.Println(f.name, "在进攻")
}
func (f *Forwards) defense() {
	if f == nil {
		return
	}
	fmt.Println(f.name, "在防守")
}

func NewForwards(name string) Player {
	return &Forwards{name}
}

type Centers struct {
	name string
}

func (f *Centers) attack() {
	if f == nil {
		return
	}
	fmt.Println(f.name, "在进攻")
}
func (f *Centers) defense() {
	if f == nil {
		return
	}
	fmt.Println(f.name, "在防守")
}

func NewCenter(name string) Player {
	return &Centers{name}
}

type ForeignCenter struct {
	name string
}

func (f *ForeignCenter) attack(what string) {
	if f == nil {
		return
	}
	fmt.Println(f.name, "在进攻")
}
func (f *ForeignCenter) defense() {
	if f == nil {
		return
	}
	fmt.Println(f.name, "在防守")
}

type Translator struct {
	f ForeignCenter
}

// 这是用户想要的接口
func (t *Translator) attack() {
	if t == nil {
		return
	}
	t.f.attack("进攻")
}
func (t *Translator) defense() {
	if t == nil {
		return
	}
	t.f.defense()
}

func NewTranslator(name string) Player {
	return &Translator{ForeignCenter{name}}
}
