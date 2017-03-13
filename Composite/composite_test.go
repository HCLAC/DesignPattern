package composite

import (
	"fmt"
	"testing"
)

func TestComponent(t *testing.T) {
	root := NewConcreateCompany("北京总公司")
	root.add(NewHRDepartment("总公司人力资源部"))
	root.add(NewFinanceDepartment("总公司财务部"))

	compA := NewConcreateCompany("上海华东分公司")
	compA.add(NewHRDepartment("上海华东分公司人力资源部"))
	compA.add(NewFinanceDepartment("上海华东分公司财务部"))
	root.add(compA)

	compB := NewConcreateCompany("南京办事处")
	compB.add(NewHRDepartment("南京办事处人力资源部"))
	compB.add(NewFinanceDepartment("南京办事处财务部"))
	compA.add(compB)

	compC := NewConcreateCompany("杭州办事处")
	compC.add(NewHRDepartment("杭州办事处人力资源部"))
	compC.add(NewFinanceDepartment("杭州办事处财务部"))
	compA.add(compC)

	fmt.Println("结构体：")
	root.display(1)

	fmt.Println("职责：")
	root.lineOfDuty()
}
