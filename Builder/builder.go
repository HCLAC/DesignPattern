/*
 Builder 生成器模式：
        （建造者模式）将一个复杂对象的构建与它表示分离，使得同样的构建过程可以创建不同的表示
 个人想法：建造的过程是一定的，建造时的细节交给具体产品来实现，在指挥者中建造，在建造者中实现细节
 作者：   HCLAC
 日期：   20170306
*/

package builder

import (
	"fmt"
)

type IBuilder interface {
	head()
	body()
	foot()
	hand()
}
type Thin struct {
}

func (t *Thin) head() {
	fmt.Println("我的头很瘦")
}

func (t *Thin) body() {
	fmt.Println("我的身体很瘦")
}
func (t *Thin) foot() {
	fmt.Println("我的脚很瘦")
}
func (t *Thin) hand() {
	fmt.Println("我的身体手很瘦")
}

type Fat struct {
}

func (t *Fat) head() {
	fmt.Println("我的头很胖")
}

func (t *Fat) body() {
	fmt.Println("我的身体很胖")
}
func (t *Fat) foot() {
	fmt.Println("我的脚很胖")
}
func (t *Fat) hand() {
	fmt.Println("我的身体手很胖")
}

type Director struct {
	person IBuilder
}

func (d *Director) CreatePerson() {
	if d == nil {
		return
	}
	d.person.head()
	d.person.body()
	d.person.foot()
	d.person.hand()
}
