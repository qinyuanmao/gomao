package logger

import (
	"fmt"
	"strings"
	"time"

	"github.com/samber/lo"
)

func Error(messages ...any) {
	New(1).Error(messages...)
}
func Errorf(message string, args ...any) {
	New(1).Errorf(message, args...)
}

func Debug(messages ...any) {
	New(1).Debug(messages...)
}
func Debugf(message string, args ...any) {
	New(1).Debugf(message, args...)
}

func Info(messages ...any) {
	New(1).Info(messages...)
}
func Infof(message string, args ...any) {
	New(1).Infof(message, args...)
}

func Warning(messages ...any) {
	New(1).Warning(messages...)
}
func Warningf(message string, args ...any) {
	New(1).Warningf(message, args...)
}

func Panic(messages ...any) {
	New(1).Panic(messages...)
}
func Panicf(message string, args ...any) {
	New(1).Panicf(message, args...)
}

func getMessage(status, caller string, messages ...any) string {
	messageArray := lo.Map(messages, func(message any, _ int) string {
		return fmt.Sprintf("[%s]%s\n%s\n%v", status, time.Now().Format("2006-01-02 15:04:05"), caller, message)
	})
	return strings.Join(messageArray, "\n")
}
