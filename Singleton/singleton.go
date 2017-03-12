/*
 Singleton 单件：
        保证一个类仅有一个实例，并提供一个访问它的全局访问点
 个人想法：用Go实现时，巧妙使用包级别的变量声明规则：小写字母的包级别变量是不对外开放的，
         创建实例时，用同步库sync.Once来保证全局只有一个对象实例。
 作者：   HCLAC
 日期：   20170305
*/
package singleton

import (
	"fmt"
	"sync"
)

// 全局实例者
type singleton struct {
	data int
}

// 定义一个包级别的private实例变量
var sin *singleton

// 同步Once,保证每次调用时，只有第一次生效
var once sync.Once

// 获取实例对象函数
func GetSingleton() *singleton {
	once.Do(func() {
		sin = &singleton{12}
	})
	fmt.Println("实例对象的信息和地址", sin, &sin)
	return sin
}
