package main

import "fmt"

type Company interface {
	Running()
}

type BigCompany struct {
	Worker
}

func (this *BigCompany) Running() {
	fmt.Printf("员工都是螺丝钉\n")
	this.Worker.Leaving()
	fmt.Printf("员工跑路\n")
}

type SmallCompany struct {
	Worker
}

func (this *SmallCompany) Running() {
	fmt.Printf("随便一个员工都是骨干\n")
	this.Worker.Leaving()
	fmt.Printf("员工跑路\n")
}

type Worker interface {
	Leaving()
}
type GoodWorker struct {
}

func (this *GoodWorker) Leaving() {
	fmt.Printf("好员工跑路")
}

type NormalWorker struct {
}

func (this *NormalWorker) Leaving() {
	fmt.Printf("普通员工跑路")
}

type DrawApi interface {
	DrawCircle(radius int, x int, y int)
}

type RedCircle struct {
}

func (this *RedCircle) DrawCircle(radius int, x int, y int) {
	fmt.Printf("radius:%v x:%v y:%v\n", radius, x, y)
}

type GreenCircle struct {
}

func (this *GreenCircle) DrawCircle(radius int, x int, y int) {
	fmt.Printf("radius:%v x:%v y:%v\n", radius, x, y)
}

type Shape struct {
	drawApi DrawApi
}

func (this *Shape) Shape(api DrawApi) {
	this.drawApi = api
}

type Circle struct {
	shape  Shape
	x      int
	y      int
	radius int
}

func main() {
	pgoodWorker := &GoodWorker{}
	pnormalWorker := &NormalWorker{}
	pBigCompany := &BigCompany{Worker: pgoodWorker}
	pBigCompany.Running()

	pBigCompany2 := &BigCompany{Worker: pnormalWorker}
	pBigCompany2.Running()

	psmallCompany := &SmallCompany{Worker: pgoodWorker}
	psmallCompany.Running()

	psmallCompany2 := &SmallCompany{Worker: pnormalWorker}
	psmallCompany2.Running()

	return
}
