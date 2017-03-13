/*
   Template Methed模板方法：
        定义一个操作中的算法的骨架，而将一些具体步骤延迟到子类中。
		模板方法使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。
 个人想法：与建造者：一个是行为型模式，一个是创建型模式
 作者：   HCLAC
 日期：   20170309
*/
package template

import (
	"fmt"
)

type getfood interface {
	first()
	secend()
	three()
}

type template struct {
	g getfood
}

func (b *template) getsomefood() {
	if b == nil {
		return
	}
	b.g.first()
	b.g.secend()
	b.g.three()
}

type bingA struct {
	template
}

func NewBingA() *bingA {
	b := bingA{}
	return &b
}

func (b *bingA) first() {
	if b == nil {
		return
	}
	fmt.Println("打开冰箱")
}

func (b *bingA) secend() {
	if b == nil {
		return
	}
	fmt.Println("拿出东西")
}

func (b *bingA) three() {
	if b == nil {
		return
	}
	fmt.Println("关闭冰箱")
}

type Guo struct {
	template
}

func NewGuo() *Guo {
	b := Guo{}
	return &b
}

func (b *Guo) first() {
	if b == nil {
		return
	}
	fmt.Println("打开锅")
}

func (b *Guo) secend() {
	if b == nil {
		return
	}
	fmt.Println("拿出东西锅")
}

func (b *Guo) three() {
	if b == nil {
		return
	}
	fmt.Println("关闭锅")
}
