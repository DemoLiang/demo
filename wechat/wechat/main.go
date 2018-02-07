package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"sync"
)

var conf Conf
var AccessToken string
var WeChatAuthLock *sync.RWMutex

//var logger *func()

func main() {
	logger := GetDefaultLogger()
	fmt.Printf("%T", logger)
	cfgFile := flag.String("conf", "", "config file path")
	flag.Parse()

	if cfgFile == nil || *cfgFile == "" {
		flag.Usage()
		return
	}
	cfg, err := os.Open(*cfgFile)
	if err != nil {
		panic(err)
	}
	defer cfg.Close()

	dec := json.NewDecoder(cfg)
	err = dec.Decode(&conf)
	if err != nil {
		panic(err)
	}

	logger.info("%v", conf)

	InitBeego()
	AddRouter()
	InitWeChatAccessToken()
	go UpdateWeChatAccessToken()
	beego.Run()

}
