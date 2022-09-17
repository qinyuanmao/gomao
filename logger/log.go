package logger

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/thoas/go-funk"
)

func Error(messages ...any) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.Red.Println(getMessage("ERROR", f.Name(), line, messages))
	}
}
func Errorf(message string, args ...any) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.Red.Println(getMessage("ERROR", f.Name(), line, fmt.Sprintf(message, args...)))
	}
}

func Debug(messages ...any) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.Green.Println(getMessage("DEBUG", f.Name(), line, messages))
	}
}
func Debugf(message string, args ...any) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.Green.Println(getMessage("DEBUG", f.Name(), line, fmt.Sprintf(message, args...)))
	}
}

func Info(messages ...any) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.White.Println(getMessage("INFO", f.Name(), line, messages))
	}
}
func Infof(message string, args ...any) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.White.Println(getMessage("INFO", f.Name(), line, fmt.Sprintf(message, args...)))
	}
}

func Warning(messages ...any) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.Yellow.Println(getMessage("WARNING", f.Name(), line, messages))
	}
}
func Warningf(message string, args ...any) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.Yellow.Println(getMessage("WARNING", f.Name(), line, fmt.Sprintf(message, args...)))
	}
}

func Panic(messages ...any) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.Red.Println(getMessage("PANIC", f.Name(), line, messages))
	}
	panic(getMessage("PANIC", f.Name(), line, messages))
}
func Panicf(message string, args ...any) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.Red.Println(getMessage("PANIC", f.Name(), line, fmt.Sprintf(message, args...)))
	}
	panic(getMessage("PANIC", f.Name(), line, fmt.Sprintf(message, args...)))
}

func getMessage(status, fileName string, line int, messages ...any) string {
	messageArray := funk.Map(messages, func(message any) string {
		return fmt.Sprintf("[%s][%s][%s:%d]%v", status, time.Now().Format("2006-01-02 15:04:05"), fileName, line, message)
	}).([]string)
	return strings.Join(messageArray, "\n")
}
