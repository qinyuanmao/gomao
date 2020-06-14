package logger

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/thoas/go-funk"
)

func Error(messages ...interface{}) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.Red.Println(getMessage("ERROR", f.Name(), line, messages))
	}
}
func Errorf(message string, args ...interface{}) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.Red.Println(getMessage("ERROR", f.Name(), line, fmt.Sprintf(message, args...)))
	}
}

func Debug(messages ...interface{}) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.Green.Println(getMessage("DEBUG", f.Name(), line, messages))
	}
}
func Debugf(message string, args ...interface{}) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.Green.Println(getMessage("DEBUG", f.Name(), line, fmt.Sprintf(message, args...)))
	}
}

func Info(messages ...interface{}) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.White.Println(getMessage("INFO", f.Name(), line, messages))
	}
}
func Infof(message string, args ...interface{}) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.White.Println(getMessage("INFO", f.Name(), line, fmt.Sprintf(message, args...)))
	}
}

func Warning(messages ...interface{}) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.Yellow.Println(getMessage("WARNING", f.Name(), line, messages))
	}
}
func Warningf(message string, args ...interface{}) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		color.Yellow.Println(getMessage("WARNING", f.Name(), line, fmt.Sprintf(message, args...)))
	}
}

func getMessage(status, fileName string, line int, messages ...interface{}) string {
	messageArray := funk.Map(messages, func(message interface{}) string {
		return fmt.Sprintf("[%s][%s][%s:%d]%v", status, time.Now().Format(time.RFC3339), fileName, line, message)
	}).([]string)
	return strings.Join(messageArray, "\n")
}
