package main

import "encoding/xml"

const (
	WeChatHost = "https://api.weixin.qq.com"
)

const (
	ReqPathError = iota
	GetToken
	TempLateSend
	MenuCreate
	GetOpenIdByAuth2
)

var WeChatReqPath = map[int]string{
	GetToken:         "/cgi-bin/token",
	TempLateSend:     "/cgi-bin/message/template/send",
	MenuCreate:       "/cgi-bin/menu/create",
	GetOpenIdByAuth2: "/sns/oauth2/access_token",
}

const (
	ReqMessageTypeText     = "text"
	ReqMessageTypeImage    = "image"
	ReqMessageTypeLink     = "link"
	ReqMessageTypeLocation = "location"
	ReqMessageTypeVoice    = "voice"
	ReqMessageTypeEvent    = "event"

	RespMessageTypeText  = ReqMessageTypeText
	RespMessageTypeMusic = "music"
	RespMessageTypeNews  = "news"

	EventTypeSubscribe   = "subscribe"
	EventTypeUnSubscribe = "unsubscribe"
	EventTypeClick       = "CLICK"
	EventTypeView        = "VIEW"
)

const (
	MenuBT01Key = "BT01"
	MenuBT02Key = "BT02"
	MenuBT03Key = "BT03"
)

type WeChatConf struct {
	AppId     string
	AppSecret string
}

type HttpConf struct {
	LogPrefix      string
	HttpPort       int
	HttpsPort      int
	Host           string
	ServerKey      string
	ServerCrt      string
	WebConfProfile string
}

type AppConf struct {
	Http HttpConf
}

type Conf struct {
	App    AppConf
	WeChat WeChatConf
}

type WeChatAuthOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}

type WeChatBaseMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Event        string   `xml:"Event"`
}

type WeChatTextMessage struct {
	WeChatBaseMessage
	Content string `xml:"Content"`
}

type WeChatMenuClick struct {
	WeChatBaseMessage
	Event    string `xml:"Event"`
	EventKey string `xml:"EventKey"`
}
