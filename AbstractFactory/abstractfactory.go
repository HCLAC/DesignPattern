/*
 Abstract Factory 抽象工厂模式：
          提供一个创建一系列相关或者相互依赖对象的接口，而无需指定他们具体的类
 个人想法：工厂模式和抽象工厂模式：感觉抽象工厂可以叫集团模式，工厂模式下，
         是一个工厂下，对产品的每一个具体生成分配不同的流水线；
         集团模式：在集团下，有不同的工厂，可以生成不同的产品，
         每个工厂生产出来的同一个型号产品具体细节是不一样
 作者：   HCLAC
 日期：   20170306
*/
package abstractfactory

import (
	"fmt"
)

type User struct {
	id   int
	name string
}

func (u *User) Id() int {
	if u == nil {
		return -1
	}
	return u.id
}

func (u *User) SetId(id int) {
	if u == nil {
		return
	}
	u.id = id
}

func (u *User) Name() string {
	if u == nil {
		return ""
	}
	return u.name
}

func (u *User) SetName(name string) {
	if u == nil {
		return
	}
	u.name = name
}

type Department struct {
	id   int
	name string
}

func (d *Department) Id() int {
	if d == nil {
		return -1
	}
	return d.id
}
func (d *Department) SetId(id int) {
	if d == nil {
		return
	}
	d.id = id
}
func (d *Department) Name() string {
	if d == nil {
		return ""
	}
	return d.name
}
func (d *Department) SetName(name string) {
	if d == nil {
		return
	}
	d.name = name
}

type IUser interface {
	insert(*User)
	getUser(int) *User
}

type SqlServerUser struct {
}

func (s *SqlServerUser) insert(u *User) {
	if s == nil {
		return
	}
	fmt.Println("往SqlServer的User表中插入一条User", u)
}

func (s *SqlServerUser) getUser(id int) (u *User) {
	if s == nil {
		return nil
	}
	u = &User{id, "hclacS"}
	fmt.Println("从SqlServer的User表中获取一条User", *u)
	return
}

type AccessUser struct {
}

func (s *AccessUser) insert(u *User) {
	if s == nil {
		return
	}
	fmt.Println("往AccessUser的User表中插入一条User", *u)
}

func (s *AccessUser) getUser(id int) (u *User) {
	if s == nil {
		return nil
	}
	u = &User{id, "hclacA"}
	fmt.Println("从AccessUser的User表中获取一条User", *u)
	return
}

type IDepartment interface {
	insert(*Department)
	getDepartment(int) *Department
}

type SqlServerDepartment struct {
}

func (s *SqlServerDepartment) insert(d *Department) {
	if s == nil {
		return
	}
	fmt.Println("往SqlServer的Department表中插入一条Department", *d)
}

func (s *SqlServerDepartment) getDepartment(id int) (u *Department) {
	if s == nil {
		return nil
	}
	u = &Department{id, "hclacDS"}
	fmt.Println("从SqlServer的Department表中获取一条Department", *u)
	return
}

type AccessDepartment struct {
}

func (s *AccessDepartment) insert(u *Department) {
	if s == nil {
		return
	}
	fmt.Println("往AccessDepartment的Department表中插入一条Department", *u)
}

func (s *AccessDepartment) getDepartment(id int) (u *Department) {
	if s == nil {
		return nil
	}
	u = &Department{id, "hclacDA"}
	fmt.Println("从AccessDepartment的Department表中获取一条Department", *u)
	return
}

type Ifactory interface {
	createUser() IUser
	createDepartment() IDepartment
}

type SqlServerFactory struct {
}

func (s *SqlServerFactory) createUser() IUser {
	if s == nil {
		return nil
	}
	u := &SqlServerUser{}
	return u
}

func (s *SqlServerFactory) createDepartment() IDepartment {
	if s == nil {
		return nil
	}
	u := &SqlServerDepartment{}
	return u
}

type AccessFactory struct {
}

func (s *AccessFactory) createUser() IUser {
	if s == nil {
		return nil
	}
	u := &AccessUser{}
	return u
}

func (s *AccessFactory) createDepartment() IDepartment {
	if s == nil {
		return nil
	}
	u := &AccessDepartment{}
	return u
}

type DataAccess struct {
	db string
}

func (d *DataAccess) createUser(db string) IUser {
	if d == nil {
		return nil
	}

	var u IUser

	if db == "sqlserver" {
		u = new(SqlServerUser)
	} else if db == "access" {
		u = new(AccessUser)
	}
	return u
}

func (d *DataAccess) createDepartment(db string) IDepartment {
	if d == nil {
		return nil
	}

	var u IDepartment

	if db == "sqlserver" {
		u = new(SqlServerDepartment)
	} else if db == "access" {
		u = new(AccessDepartment)
	}
	return u
}
