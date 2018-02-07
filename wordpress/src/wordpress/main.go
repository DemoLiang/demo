package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"golib"
)

var conf Conf
var logger *golib.Log
var version string

type HttpConf struct {
	LogPrefix   string
	Host     string
	Port     int
	WebConfPath string
	StaticFile  string
	HttpsPort   int
	HttpPort    int
	DB          bool
}
type AppConf struct {
	Http HttpConf
}
type Conf struct {
	Base golib.BaseConf
	App AppConf
}


func enableHTTPS() {
	beego.BConfig.Listen.HTTPSCertFile = "cert/jfg_ser.crt"
	beego.BConfig.Listen.HTTPSKeyFile = "cert/jfg_ser.key"
	beego.BConfig.Listen.HTTPSPort = conf.App.Http.HttpsPort
	beego.BConfig.Listen.EnableHTTPS = true
}

func InitBeego() {
	enableHTTPS()
	beego.LoadAppConfig("ini", conf.App.Http.WebConfPath)
}

func main() {
	logger = golib.GetDefaultLogger()

	cfgFile := flag.String("conf", "../../conf/wordpress.json", "config file path")
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

	bs, err := ioutil.ReadAll(cfg)
	if err != nil {
		panic(err)
	}

	// decode twice here, get port, setenv, replace host,port in json
	dec := json.NewDecoder(bytes.NewReader(bs))
	err = dec.Decode(&conf)
	if err != nil {
		panic(err)
	}

	conf.Base.Log.Prefix = conf.App.Http.LogPrefix
	logger.ApplyConf(conf.Base.Log)

	logger.Info("%s version:%s startup", os.Args[0], version)
	logger.Info("use config file:%s", *cfgFile)

	//// Init DB env
	//if conf.App.Http.DB {
	//	golib.DB.InitConn(&conf.Base.MySQL, 5, 10)
	//	golib.InitService(&conf.Base, golib.ServiceDB)
	//}
	//golib.InitService(&conf.Base, golib.ServiceOSS|golib.ServiceRedis)

	InitBeego()

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	beego.Run()
}
