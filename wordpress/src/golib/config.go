package golib


// LogServerConf 日志配置

type LogServerConf struct {
	Level  string
	Prefix string
	Syslog int
	IP     string
	Port   int
}

// mysql conf
type MysqlConf struct {
	IP        string
	Port      int
	DB        string
	User      string
	Password  string
	MaxIdle   int
	MaxActive int
}

// BaseConf APP通用配置
type BaseConf struct {
	Log       LogServerConf
}