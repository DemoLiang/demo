package main

import "github.com/astaxie/beego"

func AddRouter() {
	beego.Router("/upload", &UploadController{})
	beego.Router("/*", &MainController{})
}
