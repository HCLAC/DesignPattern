/*
 Proxy 代理模式：
        为其他对象提供一种代理，以控制对这个对象的访问。
 个人想法：代理和代理的对象接口一致，客户端不知道代理对象，
        而适配器是客户端想要适配器的接口，适配器对象的接口和客户端想要的不一样，
        适配器将适配器对象的接口封装一下，改成客户端想要的接口；
 作者：   HCLAC
 日期：   20170311
*/
package proxy

import (
	"fmt"
)

type GiveGift interface {
	giveDolls()
	giveFlowers()
	giveChocolate()
}

type Girl struct {
	name string
}

func (g *Girl) Name() string {
	if g == nil {
		return ""
	}
	return g.name
}

func (g *Girl) SetName(name string) {
	if g == nil {
		return
	}
	g.name = name
}

type Pursuit struct {
	girl Girl
}

func (p *Pursuit) giveDolls() {
	if p == nil {
		return
	}
	fmt.Println(p.girl.name, "送你洋娃娃")
}

func (p *Pursuit) giveFlowers() {
	if p == nil {
		return
	}
	fmt.Println(p.girl.name, "送你玫瑰花")
}

func (p *Pursuit) giveChocolate() {
	if p == nil {
		return
	}
	fmt.Println(p.girl.name, "送你巧克力")
}

type Proxy struct {
	p Pursuit
}

func (p *Proxy) giveDolls() {
	if p == nil {
		return
	}
	p.p.giveDolls()
}

func (p *Proxy) giveFlowers() {
	if p == nil {
		return
	}
	p.p.giveFlowers()
}

func (p *Proxy) giveChocolate() {
	if p == nil {
		return
	}
	p.p.giveChocolate()
}

func NewProxy(mm Girl) *Proxy {
	gg := Pursuit{mm}
	return &Proxy{gg}
}
