package main

import (
	"log"
	"strings"
)

type ICompany interface {
	Add(ic ICompany)
	Remove(ic ICompany)
	Display(depth int)
	LineOfDuty()
}

type ConcreteCompany struct {
	Name string
	List map[ICompany]ICompany
}

func NewConcreteCompany(name string) *ConcreteCompany {
	list := make(map[ICompany]ICompany)
	return &ConcreteCompany{
		Name: name,
		List: list,
	}
}

func (cc *ConcreteCompany) Add(ic ICompany) {
	cc.List[ic] = ic
}

func (cc *ConcreteCompany) Remove(ic ICompany) {
	delete(cc.List, ic)
}

func (cc *ConcreteCompany) Display(depth int) {
	log.Println(strings.Repeat("-", depth), " ", cc.Name)
	for _, ccc := range cc.List {
		ccc.Display(depth + 2)
	}
}

func (cc *ConcreteCompany) LineOfDuty() {
	for _, ccc := range cc.List {
		ccc.LineOfDuty()
	}
}

type HRDepartment struct {
	Name string
}

func NewHRDepartment(name string) *HRDepartment {
	return &HRDepartment{
		Name: name,
	}
}

func (hrd *HRDepartment) Add(ic ICompany) {
}

func (hrd *HRDepartment) Remove(ic ICompany) {
}

func (hrd *HRDepartment) Display(depth int) {
	log.Println(strings.Repeat("-", depth), " ", hrd.Name)
}

func (hrd *HRDepartment) LineOfDuty() {
	log.Printf("%v 员工管理招聘培训管理\n", hrd.Name)
}

type FinanceDepartment struct {
	Name string
}

func NewFinanceDepartment(name string) *FinanceDepartment {
	return &FinanceDepartment{
		Name: name,
	}
}

func (fd *FinanceDepartment) Add(ic ICompany) {
}
func (fd *FinanceDepartment) Remove(ic ICompany) {
}
func (fd *FinanceDepartment) Display(depth int) {
	log.Println(strings.Repeat("-", depth), " ", fd.Name)
}
func (fd *FinanceDepartment) LineOfDuty() {
	log.Printf("%v 公司财务收支管理\n", fd.Name)
}

func main() {

	root := NewConcreteCompany("北京总公司")
	root.Add(NewHRDepartment("总公司人力资源部"))
	root.Add(NewFinanceDepartment("总公司财务"))

	com := NewConcreteCompany("上海华东分公司")
	com.Add(NewHRDepartment("上海华东分公司人力资源部"))
	com.Add(NewFinanceDepartment("上海华东分公司财务"))
	root.Add(com)

	com1 := NewConcreteCompany("南京办事处")
	com1.Add(NewHRDepartment("南京办事处人力资源部"))
	com1.Add(NewFinanceDepartment("南京办事处财务"))
	com.Add(com1)

	com2 := NewConcreteCompany("杭州办事处")
	com2.Add(NewHRDepartment("杭州办事处人力资源部"))
	com2.Add(NewFinanceDepartment("杭州办事处财务"))
	com.Add(com2)

	root.Display(1)
	root.LineOfDuty()

	return
}
