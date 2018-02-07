package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

func enableHttps() {
	beego.BConfig.Listen.HTTPSCertFile = conf.App.Http.ServerCrt
	beego.BConfig.Listen.HTTPSKeyFile = conf.App.Http.ServerKey
	beego.BConfig.Listen.HTTPSPort = conf.App.Http.HttpsPort
	beego.BConfig.Listen.HTTPPort = conf.App.Http.HttpPort
	beego.BConfig.Listen.HTTPAddr = conf.App.Http.Host
	beego.BConfig.Listen.EnableHTTPS = false
	beego.BConfig.Listen.EnableHTTP = true
	beego.BConfig.CopyRequestBody = true
}

func InitBeego() {
	enableHttps()
	beego.LoadAppConfig("ini", conf.App.Http.WebConfProfile)
}

func WeChatGetAuth(conf WeChatConf) string {
	value := url.Values{}

	value.Add("grant_type", "client_credential")
	value.Add("appid", conf.AppId)
	value.Add("secret", conf.AppSecret)

	body := strings.NewReader(value.Encode())
	res, err := http.Post(WeChatHost+WeChatReqPath[GetToken], "application/x-www-form-urlencoded", body)
	if err != nil {
		fmt.Printf("%v", err)
	}

	defer res.Body.Close()
	var jsonObj WeChatAuthOutPut
	jsonDecoder := json.NewDecoder(res.Body)
	err = jsonDecoder.Decode(&jsonObj)
	logger.info("Get Access_token:%s", jsonObj)
	return jsonObj.AccessToken
}

func InitWeChatAccessToken() {
	WeChatAuthLock = new(sync.RWMutex)
	AccessToken = WeChatGetAuth(conf.WeChat)
}

func UpdateWeChatAccessToken() {
	for {
		time.Sleep(time.Minute * 100)
		WeChatAuthLock.Lock()
		AccessToken = WeChatGetAuth(conf.WeChat)
		WeChatAuthLock.Unlock()
	}

}

func PushJsonConten(act int, content string) (resp string, err error) {

	headData := url.Values{}
	headData.Set("access_token", AccessToken)
	u, _ := url.ParseRequestURI(WeChatHost)
	u.Path = WeChatReqPath[act]
	u.RawQuery = headData.Encode()
	urlStr := fmt.Sprintf("%v", u)
	logger.info("request url-->%v", urlStr)

	client := &http.Client{}
	request, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(content))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	res, err := client.Do(request)
	if err != nil {
		logger.info("post msg to wechat error:%v", err)
		return "", err
	}
	result, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		logger.info("read result error:%v", err)
		return "", err
	}
	logger.info("post msg to wechat result:%s", result)
	return string(result), nil
}
