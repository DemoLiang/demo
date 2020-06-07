package main

import "log"

type State interface {
	WriteProgram(work Work)
}
type Work struct {
	hour int
	current State
	finish bool
}
func(w *Work)SetState(s State){
	w.current = s
}
func (w *Work)SetHour(hour int){
	w.hour = hour
}
func(w *Work)SetFinishState(finish bool){
	w.finish = finish
}

func (w Work)WriteProgram(){
	w.current.WriteProgram(w)
}

type ForenoonState struct {
}

func (fs ForenoonState)WriteProgram(work Work){
	if work.hour<12{
		log.Printf("上午\n")
	}else{
		work.SetState(NoonState{})
		work.WriteProgram()
	}
}

type NoonState struct {
}

func (ns NoonState)WriteProgram(work Work){
	if work.hour<13{
		log.Printf("中午\n")
	}else{
		work.SetState(AfternoonState{})
		work.WriteProgram()
	}
}
type AfternoonState struct {
}

func (as AfternoonState)WriteProgram(work Work){
	if work.hour<17{
		log.Printf("下午\n")
	}else{
		work.SetState(EveningState{})
		work.WriteProgram()
	}
}

type EveningState struct {
}
func(es EveningState)WriteProgram(work Work){
	if work.finish{
		work.SetState(RestState{})
		work.WriteProgram()
	}else{
		if work.hour<21{
			log.Printf("晚间\n")
		}else{
			work.SetState(SleepingState{})
			work.WriteProgram()
		}
	}
}

type SleepingState struct {
}

func(ss SleepingState)WriteProgram(work Work){
	log.Printf("睡着了\n")
}

type RestState struct {
}
func(rs RestState)WriteProgram(work Work){
	log.Printf("下班回家\n")
}

func main(){
	tWork:=Work{}
	tState:=ForenoonState{}
	tWork.SetState(tState)
	tWork.SetFinishState(true)
	tWork.SetHour(22)
	tWork.WriteProgram()
}