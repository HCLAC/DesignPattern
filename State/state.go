/*
 State 状态模式：
        当一个对象的内在状态改变时，允许改变其行为，这个对象看起来像是改变了其类
 个人想法：UML图很相似，策略模式是用在对多个做同样事情（统一接口）的类对象的选择上，
         而状态模式是：将对某个事情的处理过程抽象成接口和实现类的形式，
		由context保存一份state，在state实现类处理事情时，修改状态传递给context，
		由context继续传递到下一个状态处理中，
 作者：   HCLAC
 日期：   20170309
*/
package state

import (
	"fmt"
)

// 工作类 --context
type Work struct {
	hour  int
	state State
}

func (w *Work) Hour() int {
	if w == nil {
		return -1
	}
	return w.hour
}

func (w *Work) State() State {
	if w == nil {
		return nil
	}
	return w.state
}

func (w *Work) SetHour(h int) {
	if w == nil {
		return
	}
	w.hour = h
}

func (w *Work) SetState(s State) {
	if w == nil {
		return
	}
	w.state = s
}

func (w *Work) writeProgram() {
	if w == nil {
		return
	}
	w.state.writeProgram(w)
}

func NewWork() *Work {
	state := new(moringState)
	return &Work{state: state}
}

type State interface {
	writeProgram(w *Work)
}

// 上午时分状态类
type moringState struct {
}

func (m *moringState) writeProgram(w *Work) {
	if w.Hour() < 12 {
		fmt.Println("现在是上午时分", w.Hour())
	} else {
		w.SetState(new(NoonState))
		w.writeProgram()
	}
}

// 中午时分状态类
type NoonState struct {
}

func (m *NoonState) writeProgram(w *Work) {
	if w.Hour() < 13 {
		fmt.Println("现在是中午时分", w.Hour())
	} else {
		w.SetState(new(AfternoonState))
		w.writeProgram()
	}
}

// 下午时分状态类
type AfternoonState struct {
}

func (m *AfternoonState) writeProgram(w *Work) {
	if w.Hour() < 17 {
		fmt.Println("现在是下午时分", w.Hour())
	} else {
		w.SetState(new(EveningState))
		w.writeProgram()
	}
}

// 晚上时分状态类
type EveningState struct {
}

func (m *EveningState) writeProgram(w *Work) {
	if w.Hour() < 21 {
		fmt.Println("现在是晚上时分", w.Hour())
	} else {
		fmt.Println("现在开始睡觉", w.Hour())
	}
}
