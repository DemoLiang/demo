package main

import "log"

type UnitedNations interface {
	ForwardMessage(message string,country Country)
}

type UnitedNationsSecurityCouncil struct {
	USA
	Iraq
}

func(unsc UnitedNationsSecurityCouncil)ForwardMessage(message string,country Country){
	switch country.(type) {
	case USA:
		unsc.Iraq.GetMessage(message)
	case Iraq:
		unsc.USA.GetMessage(message)
	default:
		log.Printf("the country is not a member of usa or iraq\n")
	}
}

type Country interface {
	SendMessage(message string)
	GetMessage(message string)
}
type USA struct {
	UnitedNations
}

func (usa USA)SendMessage(message string){
	usa.UnitedNations.ForwardMessage(message,usa)
}

func (usa USA)GetMessage(message string){
	log.Printf("美国收到对方消息:%v\n",message)
}
type Iraq struct {
	UnitedNations
}

func (iraq Iraq)SendMessage(message string){
	iraq.UnitedNations.ForwardMessage(message,iraq)
}

func (iraq Iraq)GetMessage(message string){
	log.Printf("伊拉克收到对方消息：%v\n",message)
}

func main(){
	//创建一个具体的中介
	tMediator:=UnitedNationsSecurityCouncil{}
	//创建具体同事，并且让他认识中介者
	tColleageA:=USA{
		UnitedNations:tMediator,
	}
	tColleageB:=Iraq{
		UnitedNations:tMediator,
	}

	tMediator.USA = tColleageA
	tMediator.Iraq = tColleageB
	tColleageA.SendMessage("停止核武器研发，否则发动战争\n")
	tColleageB.SendMessage("我们没有研发核武器，也不怕战争\n")
	return
}