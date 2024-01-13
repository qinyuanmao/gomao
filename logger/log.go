package logger

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/samber/lo"
)

// 多层级调用，初始时默认是 3
var logger *Logger
var _once = sync.Once{}

func InitCallerLevel(level int) {
	_once.Do(func() {
		logger = newLevel(level, 3) // 3 是 getCaller, Error, Errorf... 的上一层级
	})
}

func getLogger() *Logger {
	if logger == nil {
		InitCallerLevel(1)
	}
	return logger
}

func Error(messages ...any) {
	getLogger().Error(messages...)
}
func Errorf(message string, args ...any) {
	getLogger().Errorf(message, args...)
}

func Debug(messages ...any) {
	getLogger().Debug(messages...)
}
func Debugf(message string, args ...any) {
	getLogger().Debugf(message, args...)
}

func Info(messages ...any) {
	getLogger().Info(messages...)
}
func Infof(message string, args ...any) {
	getLogger().Infof(message, args...)
}

func Warning(messages ...any) {
	getLogger().Warning(messages...)
}
func Warningf(message string, args ...any) {
	getLogger().Warningf(message, args...)
}

func Panic(messages ...any) {
	getLogger().Panic(messages...)
}
func Panicf(message string, args ...any) {
	getLogger().Panicf(message, args...)
}

func getMessage(status, caller string, messages ...any) string {
	messageArray := lo.Map(messages, func(message any, _ int) string {
		return fmt.Sprintf("[%s] %s\n%s\n\t%v", status, time.Now().Format("2006-01-02 15:04:05"), caller, message)
	})
	return strings.Join(messageArray, "\n")
}
