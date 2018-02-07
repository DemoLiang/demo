package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type LogLevel int

var logger = NewLogger(os.Args[0] + " ")

type LogServerConf struct {
	Level  string
	Prefix string
	Syslog int
}

const (
	LogError LogLevel = iota
	LogInfo
	LogDebug
)

type Log struct {
	level    LogLevel
	prefix   string
	syslog   bool
	ioWriter io.Writer

	output [LogDebug + 1]*log.Logger
	l      sync.Mutex
}

func NewLogger(prefix string) *Log {
	ret := &Log{
		level:  LogDebug,
		prefix: prefix,
		syslog: false,
	}

	for i := LogError; i <= LogDebug; i++ {
		ret.output[i] = log.New(os.Stdout, prefix, 0)
	}
	ret.SetLevel(ret.level)
	return ret
}

func GetDefaultLogger() *Log {
	return logger
}

func (p *Log) SetPrefix(prefix string) {
	p.l.Lock()
	defer p.l.Unlock()

	p.prefix = prefix
	for i := LogError; i <= LogDebug; i++ {
		p.output[i].SetPrefix(p.prefix)
	}
}

func (p *Log) SetLevel(level LogLevel) {
	p.l.Lock()
	defer p.l.Unlock()

	p.level = level
	var flag = 0

	if !p.syslog {
		flag |= (log.Ldate | log.Ltime)
	}
	if p.level == LogDebug {
		flag |= log.Lshortfile
	}
	for i := LogError; i <= LogDebug; i++ {
		if i == LogError {
			flag |= log.Lshortfile
		}
		p.output[i].SetFlags(flag)
	}
}

func (p *Log) SetLevelStr(level string) error {
	switch level {
	case "error":
		p.SetLevel(LogError)
	case "info":
		p.SetLevel(LogInfo)
	case "debug":
		p.SetLevel(LogDebug)
	default:
		return fmt.Errorf("invalid level:%s", level)
	}
	return nil
}

func (p *Log) ApplyConf(conf LogServerConf) {
	if conf.Prefix != "" {
		logger.SetPrefix(conf.Prefix)
	}
	if err := logger.SetLevelStr(conf.Level); err != nil {
		panic(err)
	}
}

func (p *Log) Errorf(format string, v ...interface{}) {
	if p.level < LogError {
		return
	}
	p.output[LogError].Output(2, fmt.Sprintf(format, v...))
}

func (p *Log) info(format string, v ...interface{}) {
	if p.level < LogInfo {
		return
	}
	p.output[LogInfo].Output(2, fmt.Sprintf(format, v...))
}
func (p *Log) Debug(format string, v ...interface{}) {
	if p.level < LogDebug {
		return
	}
	p.output[LogDebug].Output(2, fmt.Sprintf(format, v...))
}
