package main

import "fmt"

type PersonalInfo struct {
	Name string
	Sex  string
	Age  string
}

type WorExperience struct {
	TimeArea string
	Company  string
}

type Resume struct {
	PersonalInfo
	WorExperience
}

func (this *Resume) SetPersonalInfo(name, sex, age string) {
	this.Name = name
	this.Sex = sex
	this.Age = age
}

func (this *Resume) SetWorkExperience(timeArea string, company string) {
	this.TimeArea = timeArea
	this.Company = company
}

func (this *Resume) Display() {
	fmt.Printf("%v %v %v\n", this.Name, this.Sex, this.Age)
	fmt.Printf("work experience:%v %v\n", this.TimeArea, this.Company)
}

func (this *Resume) Clone() *Resume {
	resume := *this
	return &resume
}

func main() {
	r1 := &Resume{}
	r1.SetPersonalInfo("job", "man", "28")
	r1.SetWorkExperience("2014-7", "XX_company")

	r2 := r1.Clone()
	r2.SetWorkExperience("2020-1", "YY_company")

	r3 := r1.Clone()
	r3.SetPersonalInfo("tim", "man", "30")

	r1.Display()
	r2.Display()
	r3.Display()
	return
}
