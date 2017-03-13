/*
 Interpreter 解释器模式：
        给定一个语言，定义它的文法的一种表示，
		并定义一个解释器，这个解释器使用该表示来解释语言中的句子
 个人想法：
 作者：   HCLAC
 日期：   20170310
*/

package interpreter

import (
	"fmt"
)

type Context struct {
	text string
}

// 抽象表达式
type IAbstractExpression interface {
	Interpret(*Context)
}

// 终结符表达式
type TerminalExpression struct {
}

func (t *TerminalExpression) Interpret(context *Context) {
	if t == nil {
		return
	}
	context.text = context.text[:len(context.text)-1]
	fmt.Println(context)
}

// 非终结符表达式
type NonterminalExpression struct {
}

func (t *NonterminalExpression) Interpret(context *Context) {
	if t == nil {
		return
	}
	context.text = context.text[:len(context.text)-1]
	fmt.Println(context)
}
