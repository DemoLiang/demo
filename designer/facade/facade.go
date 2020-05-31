package main

import "log"

type Power struct {
}

func(p Power)PowerOn(){
	log.Printf("电源通电\n")
}

func (p Power)PowerOff(){
	log.Printf("电源关电\n")
}

type Mainboard struct {
}

func(m Mainboard)PowerOn(){
	log.Printf("主板通电\n")
}

func (m Mainboard)PowerOff(){
	log.Printf("主板关电\n")
}

type CPU struct {
}

func(cpu CPU)PowerOn(){
	log.Printf("CPU通电\n")
}

func(cpu CPU)PowerOff(){
	log.Printf("CPU关电\n")
}

type Computer interface {
	Boot()
	Shutdown()
}

type DesktopComputer struct {
	Power
	Mainboard
	CPU
}

func (dc DesktopComputer)Boot(){
	log.Printf("台式机启动\n")
	dc.Power.PowerOn()
	dc.Mainboard.PowerOn()
	dc.CPU.PowerOn()
	log.Printf("台式机启动完毕\n")
	return
}

func(dc DesktopComputer)Shutdown(){
	log.Printf("台式机关机\n")
	dc.Power.PowerOff()
	dc.Mainboard.PowerOff()
	dc.CPU.PowerOff()
	log.Printf("台式机关机完毕\n")
	return
}

type Laptop struct {
	Power
	Mainboard
	CPU
}
func(l Laptop)Boot(){
	log.Printf("笔记本电脑开机\n")
	l.Power.PowerOn()
	l.Mainboard.PowerOn()
	l.CPU.PowerOn()
	log.Printf("笔记本启动完毕\n")
}

func(l Laptop)Shutdown(){
	log.Printf("笔记本关机\n")
	l.Power.PowerOff()
	l.Mainboard.PowerOff()
	l.CPU.PowerOff()
	log.Printf("笔记本关机完毕\n")
	return
}

func NewComputer(t string)Computer{
	switch t {
	case "d":
		return DesktopComputer{}
	case "l":
		return Laptop{}
	default:
		return nil
	}
}


func main(){
	tComputer:=NewComputer("d")
	tComputer.Boot()
	log.Printf("========\n")
	tComputer.Shutdown()
	log.Printf("==========\n")
	lComputer:=NewComputer("l")
	lComputer.Boot()
	log.Printf("==========\n")
	lComputer.Shutdown()
	return
}