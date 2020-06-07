package main

import "log"

type Memento struct {
	state string
}

func(m *Memento)SetState(s string){
	m.state = s
}
func(m *Memento)GetState()string{
	return m.state
}

type Originator struct {
	state string
}

func(o *Originator)SetState(s string){
	o.state = s

}

func (o *Originator)GetState()string{
	return o.state
}

func (o *Originator)CreateMemento()*Memento{
	return &Memento{state:o.state}
}

type Caretaker struct {
	memento *Memento
}
func (c *Caretaker)GetMemento()*Memento{
	return c.memento
}
func(c *Caretaker)SetMemento(m *Memento){
	c.memento = m
}

func main(){
	o:=&Originator{state:"hello"}
	log.Printf("当前状态:%v\n",o.GetState())

	c:=new(Caretaker)
	c.SetMemento(o.CreateMemento())
	o.SetState("world")
	log.Printf("更改当前状态:%v\n",o.GetState())

	o.SetState(c.GetMemento().GetState())
	log.Printf("恢复后状态:%v\n",o.GetState())
	return
}