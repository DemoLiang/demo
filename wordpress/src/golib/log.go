package golib

import (
	"os"
	"io"
	stdlog "log"
	"sync"
	"fmt"
	"strings"
)
// LogLevel 日志级别
type LogLevel int

const (
	// LogDebug debug级别
	LogDebug LogLevel = iota
	// LogInfo info 级别
	LogInfo
	// LogError error级别
	LogError
)

// Log 日志配置结构
type Log struct {
	level    LogLevel
	prefix   string
	syslog   bool
	puWriter io.Writer
	//w [LOG_ERROR + 1]io.Writer
	output [LogError + 1]*stdlog.Logger
	l      sync.Mutex
}
// SetLevel 设置日志级别
func (p *Log) SetLevel(level LogLevel) {
	p.l.Lock()
	defer p.l.Unlock()

	p.level = level

	var flag = 0
	if !p.syslog {
		flag |= (stdlog.Ldate | stdlog.Ltime)
	}
	if p.level == LogDebug {
		flag |= stdlog.Lshortfile
	}

	for i := LogDebug; i <= LogError; i++ {
		if i == LogError {
			flag |= stdlog.Lshortfile
		}
		p.output[i].SetFlags(flag)
	}
}

// NewLogger 获取日志操作结构
func NewLogger(prefix string) *Log {
	ret := &Log{
		level:  LogDebug,
		prefix: prefix,
		syslog: false,
	}
	// default to os.Stdout
	for i := LogDebug; i <= LogError; i++ {
		// only outputs filename for debug purpose
		ret.output[i] = stdlog.New(os.Stdout, prefix, 0)
	}

	ret.SetLevel(ret.level)
	return ret
}

// SetPrefix 设置日志前缀
func (p *Log) SetPrefix(prefix string) {
	p.l.Lock()
	defer p.l.Unlock()

	p.prefix = prefix
	for i := LogDebug; i <= LogError; i++ {
		p.output[i].SetPrefix(p.prefix)
	}
}

// SetLevelStr 设置日志级别
func (p *Log) SetLevelStr(level string) error {
	if 0 == strings.Compare(level, "debug") {
		p.SetLevel(LogDebug)
	} else if 0 == strings.Compare(level, "info") {
		p.SetLevel(LogInfo)
	} else if 0 == strings.Compare(level, "error") {
		p.SetLevel(LogError)
	} else {
		return fmt.Errorf("invalid level:%s", level)
	}
	return nil
}
// ApplyConf 应用日志设置
func (p *Log) ApplyConf(conf LogServerConf) {
	if conf.Prefix != "" {
		logger.SetPrefix(conf.Prefix)
	}
	if err := logger.SetLevelStr(conf.Level); err != nil {
		panic(err)
	}

	if conf.Syslog == 1 {
		logger.Info("redirect log to %s:%d", conf.IP, conf.Port)
		p.EnableSyslog(conf.IP, conf.Port)
	}
}
// EnableSyslog windows上启动Syslog，空实现
func (p *Log) EnableSyslog(host string, port int) {
}

// Debug 输出Debug级别日志
func (p *Log) Debug(format string, v ...interface{}) {
	p.l.Lock()
	defer p.l.Unlock()

	if p.level > LogDebug {
		return
	}
	p.output[LogDebug].Output(2, fmt.Sprintf(format, v...))
}

// Info 输出info级别日志
func (p *Log) Info(format string, v ...interface{}) {
	p.l.Lock()
	defer p.l.Unlock()

	if p.level > LogInfo {
		return
	}
	p.output[LogInfo].Output(2, fmt.Sprintf(format, v...))
}

// Errorf 输出错误级别日志
func (p *Log) Errorf(format string, v ...interface{}) {
	p.l.Lock()
	defer p.l.Unlock()

	if p.level > LogError {
		return
	}
	p.output[LogError].Output(2, fmt.Sprintf(format, v...))
}

// Debug 输出Debug级别日志
func (p *Log) DebugDx(debugStr, format string, v ...interface{}) {
	p.l.Lock()
	defer p.l.Unlock()

	if p.level > LogDebug {
		return
	}

	if debugStr != "" {
		format = fmt.Sprintf("[%s] %s", debugStr, format)
	}
	p.output[LogDebug].Output(2, fmt.Sprintf(format, v...))
}

// Info 输出info级别日志
func (p *Log) InfoDx(debugStr, format string, v ...interface{}) {
	p.l.Lock()
	defer p.l.Unlock()

	if p.level > LogInfo {
		return
	}

	if debugStr != "" {
		format = fmt.Sprintf("[%s] %s", debugStr, format)
	}
	p.output[LogInfo].Output(2, fmt.Sprintf(format, v...))
}

// Info 输出info级别日志
func (p *Log) ErrorfDx(debugStr, format string, v ...interface{}) {
	p.l.Lock()
	defer p.l.Unlock()

	if p.level > LogError {
		return
	}

	if debugStr != "" {
		format = fmt.Sprintf("[%s] %s", debugStr, format)
	}
	p.output[LogError].Output(2, fmt.Sprintf(format, v...))
}

func (p *Log) LogPu(logStr string) {
	if p.puWriter != nil {
		p.puWriter.Write([]byte(logStr))
	}
}

var logger = NewLogger(os.Args[0] + " ")

// GetDefaultLogger 获取默认日志句柄
func GetDefaultLogger() *Log {
	return logger
}
