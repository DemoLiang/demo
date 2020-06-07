package main

import "log"

type Request struct {
	RequestType string
	RequestContent string
	Number int
}

type Manager interface{
	SetNext(next Manager)
	RequestHandler(request Request)
}

type CommonManager struct {
	Manager
	Name string
}

func (cm *CommonManager)SetNext(next Manager){
	cm.Manager = next
}

func (cm *CommonManager)RequestHandler(request Request){
	if request.RequestType =="请假"&&request.Number<=2{
		log.Printf("%v %v 数量 %v 已经批准\n",
			cm.Name,request.RequestContent,request.Number)
	}else{
		if cm.Manager!=nil{
			cm.Manager.RequestHandler(request)
		}
	}
}

type MajorManager struct {
	Manager
	name string
}
func(mm *MajorManager)SetNext(next Manager){
	mm.Manager = next
}

func(mm *MajorManager)RequestHandler(request Request){
	if request.RequestType=="请假"&&request.Number<=5{
		log.Printf("%v %v 数量 %v 已经批准\n",
			mm.name,request.RequestContent,request.Number)
	}else{
		if mm.Manager!=nil{
			mm.Manager.RequestHandler(request)
		}
	}
}

type GeneralManager struct {
	Manager
	name string
}

func (gm *GeneralManager)SetNext(next Manager){
	gm.Manager = next
}

func (gm *GeneralManager)RequestHandler(request Request){
	if request.RequestType == "请假"{
		log.Printf("%v %v 数量 %v 已经批准\n",
			gm.name,request.RequestContent,request.Number)
	}else if request.RequestType == "加薪"&& request.Number<500{
		log.Printf("%v %v 数量 %v 已经批准\n",
			gm.name,request.RequestContent,request.Number)
	}else if request.RequestType == "加薪" && request.Number>500{
		log.Printf("%v %v 数量 %v 再说吧\n",
			gm.name,request.RequestContent,request.Number)
	}
}

func main(){
	wenxiang:=&CommonManager{
		Name:"文祥",
	}
	xiaoyun:=&MajorManager{
		name:"晓云",
	}
	yuanlei:=&GeneralManager{
		name:"元类",
	}
	wenxiang.SetNext(xiaoyun)
	xiaoyun.SetNext(yuanlei)

	request:=Request{
		RequestType:"请假",
		RequestContent:"小菜请假",
		Number:2,
	}
	wenxiang.RequestHandler(request)
	request = Request{
		RequestType:"请假",
		RequestContent:"小菜请假",
		Number:5,
	}
	wenxiang.RequestHandler(request)
}