/*
 Factory Method 工厂方法模式：
        定义一个用于创建对象的接口，让子类决定实例化哪一个类。
        工厂方法使一个类的实例化延迟到其子类

 个人想法：简单工厂和工厂模式：简单工厂定义的是静态函数，
        一个函数处理所有的产品创建，工厂模式将创建对象过程抽象为一个类组，
		有抽象类，有对应产品的创建类，创建的过程有创建类来完成，
		工厂模式主要使用的是依赖反转原则
		（高层模块不依赖底层模块，统一依赖抽象层，抽象层不依赖细节层，细节层依赖抽象层），
		解决简单工厂的缺少开放-封闭原则
 作者：   HCLAC
 日期：   20170306
*/

package factorymethod

//"fmt"

type OperationFunc interface {
	SetNumA(float32)
	SetNumB(float32)
	Result() (float32, bool)
}

type Operation struct {
	numberA float32
	numberB float32
}

func (o *Operation) SetNumA(num float32) {
	if o == nil {
		return
	}
	o.numberA = num
}
func (o *Operation) SetNumB(num float32) {
	if o == nil {
		return
	}
	o.numberB = num
}

type OperationAdd struct {
	Operation
}

func (a *OperationAdd) Result() (float32, bool) {
	if a == nil {
		return 0, false
	}
	return a.numberA + a.numberB, true
}

type OperationSub struct {
	Operation
}

func (a *OperationSub) Result() (float32, bool) {
	if a == nil {
		return 0, false
	}
	return a.numberA - a.numberB, true
}

type CreateOperation interface {
	createoperation(string) OperationFunc
}

type COperationAdd struct {
}

func (c *COperationAdd) createoperation(op string) OperationFunc {
	if c == nil {
		return nil
	}
	return &OperationAdd{}
}

type COperationSub struct {
}

func (c *COperationSub) createoperation(op string) OperationFunc {
	if c == nil {
		return nil
	}
	return &OperationSub{}
}
