package main

import "log"

type IVisitor interface {
	Visit()
}

type ProductionVisitor struct {
}

func (v ProductionVisitor)Visit(){
	log.Printf("生成环境\n")
}

type TestingVisitor struct {
}

func (t TestingVisitor)Visit(){
	log.Printf("测试环境\n")
}

type IElement interface {
	Accept(visitor IVisitor)
}
type Element struct {
}
func (el Element)Accept(visitor IVisitor){
	visitor.Visit()
}

type EnvExample struct {
	Element
}

func(e EnvExample)Print(visitor IVisitor){
	e.Element.Accept(visitor)
}

func main(){
	e:=new(Element)
	e.Accept(new(ProductionVisitor))
	e.Accept(new(TestingVisitor))

	m:=new(EnvExample)
	m.Print(new(ProductionVisitor))
	m.Print(new(TestingVisitor))
}