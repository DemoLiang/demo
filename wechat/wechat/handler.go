package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"mime/multipart"
	"sort"
	"strings"
)

var Token string = "jfg"

type BaseController struct {
	beego.Controller
}

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	signature := this.GetString("signature")
	timestamp := this.GetString("timestamp")
	nonce := this.GetString("nonce")
	echostr := this.GetString("echostr")

	var list = []string{Token, timestamp, nonce}
	sort.Strings(list)

	sigStr := strings.Join(list, "")
	sha1 := sha1.New()
	sha1.Write([]byte(sigStr))
	hash := sha1.Sum(nil)

	logger.info("signature:%v ,timestamp:%v ,nonce:%v ,echostr:%v ,echostr:%v", signature, timestamp, nonce, echostr)
	logger.info("sigStr:%v ,list:%v ,hash:%v ", sigStr, list, hash)
	if signature == fmt.Sprintf("%x", hash) {
		logger.info("hash eq echostr!")
		this.Ctx.Output.Body([]byte(echostr))
	}
	this.Ctx.Output.Body([]byte("hello world"))
	return
}

func (this *MainController) Post() {
	var baseMsg WeChatBaseMessage
	var err error
	var respMessage string

	err = xml.Unmarshal(this.Ctx.Input.RequestBody, &baseMsg)
	if err != nil {
		logger.info("unmarshal request error:%v", err)
	}
	logger.info("%v", baseMsg)

	switch baseMsg.MsgType {
	case ReqMessageTypeText:
		var textMsg WeChatTextMessage
		err = xml.Unmarshal(this.Ctx.Input.RequestBody, &textMsg)
		if err != nil {
			logger.info("unmarshal request error:%v", err)
		}
		logger.info("%v-->%v", textMsg.FromUserName, textMsg.ToUserName)
		respMessage = "success"
	case ReqMessageTypeImage:
		logger.info("%v-->sending a image message", baseMsg.FromUserName)
		respMessage = "success"
	case ReqMessageTypeLocation:
		logger.info("%v-->sending a location message", baseMsg.FromUserName)
		respMessage = "success"
	case ReqMessageTypeLink:
		logger.info("%v-->sending a link message", baseMsg.FromUserName)
		respMessage = "success"
	case ReqMessageTypeVoice:
		logger.info("%v-->sending a voice message", baseMsg.FromUserName)
		respMessage = "success"
	case ReqMessageTypeEvent:
		switch baseMsg.Event {
		case EventTypeSubscribe:
			logger.info("%v-->sending a subscribe message", baseMsg.FromUserName)
			respMessage = "success"
		case EventTypeUnSubscribe:
			logger.info("%v-->sending a unsubscribe message", baseMsg.FromUserName)
			respMessage = "success"
		case EventTypeClick:
			logger.info("%v-->sending a click message", baseMsg.FromUserName)
			respMessage = "success"
			var menuClickMsg WeChatMenuClick
			err = xml.Unmarshal(this.Ctx.Input.RequestBody, &menuClickMsg)
			if err != nil {
				logger.info("unmarshal request error:%v", err)
			}
			switch menuClickMsg.EventKey {
			case MenuBT01Key:
				logger.info("%v-->click bt01", baseMsg.FromUserName)
				respMessage = "success"
			case MenuBT02Key:
				logger.info("%v-->click bt02", baseMsg.FromUserName)
				respMessage = "success"
			case MenuBT03Key:
				logger.info("%v-->click bt03", baseMsg.FromUserName)
				respMessage = "success"
			default:
				logger.info("%v-->click button key unknown:%v", baseMsg.FromUserName, menuClickMsg.EventKey)
				respMessage = "success"
			}
		case EventTypeView:
			logger.info("%v-->sending a view message", baseMsg.FromUserName)
			respMessage = "success"
		default:
			logger.info("%v -->send a unknown event message", baseMsg.FromUserName)
			respMessage = "success"
		}
	default:
		logger.info("%v-->send a unknown message", baseMsg.FromUserName)
		respMessage = "success"
	}

	this.Ctx.Output.Body([]byte(respMessage))
	return
}

type UploadController struct {
	BaseController
}

func (this *UploadController) Get() {
	this.TplName = "upload.html"
}
func (this *UploadController) Post() {
	filename := this.GetString("filename")
	filedName := this.GetString("filed_name")
	if filename == "" {
		fmt.Printf("filename null")
		return
	}

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	_, err := bodyWriter.CreateFormFile(filedName, filename)
	if err != nil {
		fmt.Printf("filename null")
		return
	}

	Body := bodyBuf.Bytes()
	fmt.Printf("%v", Body)
	return
}
