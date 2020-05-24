package main

import "fmt"

type Operater interface {
	Operate(int, int) int
}

type AddOperate struct {
}

func (this *AddOperate) Operate(rhs int, lhs int) int {
	return rhs + lhs
}

type MultipleOperate struct {
}

func (this *MultipleOperate) Operate(rhs int, lhs int) int {
	return rhs * lhs
}

type OperateFactory struct {
}

func NewOperateFactory() *OperateFactory {
	return &OperateFactory{}
}

func (this *OperateFactory) CreatOperate(operaterName string) Operater {
	switch operaterName {
	case "+":
		return &AddOperate{}
	case "*":
		return &MultipleOperate{}
	default:
		panic("无效运算符号")
		return nil
	}
}

func main() {
	operate := NewOperateFactory().CreatOperate("+")
	fmt.Printf("add result:%v\n", operate.Operate(1, 2))
}
