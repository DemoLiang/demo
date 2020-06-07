package main

import "log"

type Handler interface {
	HandleData(string) (interface{},error)
}

type DBHandler struct {
	DBName string
}

func(this *DBHandler)HandleData(data string)(interface{},error){
	log.Printf("get db data by query sql:%v\n",data)
	return nil,nil
}

type AppHandler struct {
	Handler *DBHandler
}

func (this *AppHandler)HandleData(data string)(interface{},error){
	if true{
		log.Printf("get data from redis\n")
	}else{
		//store data to redis
		this.Handler.HandleData(data)
	}
	return nil,nil
}

func main(){
	var iHandler Handler
	iHandler = &AppHandler{Handler:&DBHandler{DBName:"mysql"}}
	iHandler.HandleData("username")
	return
}