/*
 Strategy 策略模式：
        它定义了算法家族，分别封装起来，让它们可以相互替换，
		此模式让算法的变化，不会影响到使用算法的客户。
 个人想法：UML图很相似，策略模式是用在对多个做同样事情（统一接口）的类对象的选择上，
         而状态模式是：将对某个事情的处理过程抽象成接口和实现类的形式，
		由context保存一份state，在state实现类处理事情时，修改状态传递给context，
		由context继续传递到下一个状态处理中，
 作者：   HCLAC
 日期：   20170309
*/
package strategy

import (
	"errors"
)

type CashSuper interface {
	acceptCash(memory float32) float32
}

type CashNomal struct {
}

func (this *CashNomal) acceptCash(money float32) float32 {
	return money
}

type CashRebate struct {
	meneyRebate float32
}

func (this *CashRebate) acceptCash(money float32) float32 {
	return this.meneyRebate * money
}

type CashReturn struct {
	meneyCondition float32
	meneyReturn    float32
}

func (this *CashReturn) acceptCash(money float32) float32 {
	if money >= this.meneyCondition {
		return money - float32(int(money/this.meneyCondition*this.meneyReturn))
	} else {
		return money
	}
}

type Context struct {
	CashSuper
}

func NewCashContext(str string) (cash *Context, err error) {
	cash = new(Context)
	switch str {
	case "正常收费":
		cash.CashSuper = &CashNomal{}
	case "满300返100":
		cash.CashSuper = &CashReturn{300, 100}
	case "打8折":
		cash.CashSuper = &CashRebate{0.8}
	default:
		err = errors.New("no match")
	}
	return cash, err
}
