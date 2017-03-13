/*
 Memento 备忘录模式：
        在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。
        这样以后就可以将该对象恢复到原先保存的状态
 个人想法：将某个类的状态（某些状态，具体有该类决定）保存在另外一个类中
         （代码级别：提供一个函数能够将状态保存起来，返回出去），保存好状态的类对象是管理类的成员，
		 原来的类需要恢复时，再从管理类中获取原来的状态
 作者：   HCLAC
 日期：   20170309
*/

package memento

import (
	"fmt"
)

type GameRole struct {
	vit int
	atk int
	def int
}

func (g *GameRole) StateDisplay() {
	if g == nil {
		return
	}
	fmt.Println("角色当前状态：")
	fmt.Println("体力：", g.vit)
	fmt.Println("攻击：", g.atk)
	fmt.Println("防御：", g.def)
	fmt.Println("============")
}

func (g *GameRole) GetInitState() {
	if g == nil {
		return
	}
	g.vit = 100
	g.atk = 100
	g.def = 100
}

func (g *GameRole) Fight() {
	if g == nil {
		return
	}
	g.vit = 0
	g.atk = 0
	g.def = 0
}
func (g *GameRole) SaveState() RoleStateMemento {
	if g == nil {
		return RoleStateMemento{}
	}
	return RoleStateMemento{*g}
}

func (g *GameRole) RecoveryState(r RoleStateMemento) {
	if g == nil {
		return
	}
	g.vit = r.vit
	g.atk = r.atk
	g.def = r.def
}

type RoleStateMemento struct {
	GameRole
}

type RoleStateCaretaker struct {
	memento RoleStateMemento
}
