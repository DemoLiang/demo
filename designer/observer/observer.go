package main

import "log"

type Subject interface {
	SetState(state string)
	GetState() string
	Attach(observer Observer)
	Notify()
}

type Boss struct {
	observer []Observer
	state string
}

func (b *Boss)SetState(state string){
	b.state =state
}
func(b Boss)GetState() string{
	return b.state
}
func (b *Boss)Attach(observer Observer){
	b.observer = append(b.observer,observer)
}

func (b Boss)Notify(){
	for _,o:=range b.observer{
		o.Update()
	}
}

type Secretay struct {
	observer []Observer
	state string
}

func(s *Secretay)SetState(state string){
	s.state = state
}
func(s Secretay)GetState()string{
	return s.state
}
func(s *Secretay)Attach(observer Observer){
	s.observer = append(s.observer,observer)
}
func(s Secretay)Notify(){
	for _,o:=range s.observer{
		o.Update()
	}
}


type Observer interface {
	Update()
}

type StockObserver struct {
	name string
	subject Subject
}

func(so StockObserver)Update(){
	log.Printf("%s %s 关闭炒股软件，开始工作\n",so.subject.GetState(),so.name)
}

type NBAObserver struct {
	name string
	subject Subject
}

func(no NBAObserver)Update(){
	log.Printf("%s %s 关闭NBA直播，开始工作\n",no.subject.GetState(),no.name)
}


func NewSubject(t string)Subject{
	switch t {
	case "boss":
		return &Boss{}
	case "secretary":
		return &Secretay{}
	default:
		return nil
	}
}

func NewObserver(t,name string,subject Subject)Observer{
	switch t {
	case "stock":
		return &StockObserver{
			name:name,
			subject:subject,
		}
	case "nba":
		return &NBAObserver{
			name:name,
			subject:subject,
		}
	default:
		return nil
	}
}

func main(){
	//通知者
	tSubject:=NewSubject("boss")

	//观察者
	tObserverA:=NewObserver("stock","张三",tSubject)
	tObserverB:=NewObserver("nba","李四",tSubject)

	//通知者登记观察者
	tSubject.Attach(tObserverA)
	tSubject.Attach(tObserverB)

	//通知者状态改变
	tSubject.SetState("老板来了")

	//通知者开始通知观察者
	tSubject.Notify()
}