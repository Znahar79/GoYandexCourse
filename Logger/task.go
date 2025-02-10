package main

import (
	"log"
	"os"
)

type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
)

func (l LogLevel) IsValid() bool {
	switch l {
	case LogLevelInfo, LogLevelWarning, LogLevelError:
		return true
	default:
		return false
	}
}

type LogExtended struct {
	*log.Logger
	logLevel LogLevel // LogLevel это enum
}

func NewLogExtended() *LogExtended {
	return &LogExtended{
		Logger:   log.New(os.Stderr, "", log.LstdFlags),
		logLevel: LogLevelError,
	}
}

func (l *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
	// игнорируем сообщения, если уровень логгера меньше scrLogLvl

	if l.logLevel >= srcLogLvl {
		l.Logger.Println(prefix + msg)
	}
}

func (l *LogExtended) SetLogLevel(logLvl LogLevel) {
	l.logLevel = logLvl
}

func (l LogExtended) Infoln(msg string) {
	l.println(LogLevelInfo, "", msg)
}

func (l LogExtended) Warnln(msg string) {
	l.println(LogLevelWarning, "", msg)
}

func (l LogExtended) Errorln(msg string) {
	l.println(LogLevelError, "", msg)
}

func main() {
	logger := NewLogExtended()
	logger.SetLogLevel(LogLevelWarning)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")
}
