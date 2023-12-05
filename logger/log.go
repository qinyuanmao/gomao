package logger

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/samber/lo"
)

func Error(messages ...any) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		fileName, _ := f.FileLine(pc)
		color.Red.Println(getMessage("ERROR", fileName, line, messages))
	}
}
func Errorf(message string, args ...any) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		fileName, _ := f.FileLine(pc)
		color.Red.Println(getMessage("ERROR", fileName, line, fmt.Sprintf(message, args...)))
	}
}

func Debug(messages ...any) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		fileName, _ := f.FileLine(pc)
		color.Green.Println(getMessage("DEBUG", fileName, line, messages))
	}
}
func Debugf(message string, args ...any) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		fileName, _ := f.FileLine(pc)
		color.Green.Println(getMessage("DEBUG", fileName, line, fmt.Sprintf(message, args...)))
	}
}

func Info(messages ...any) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		fileName, _ := f.FileLine(pc)
		color.White.Println(getMessage("INFO", fileName, line, messages))
	}
}
func Infof(message string, args ...any) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		fileName, _ := f.FileLine(pc)
		color.White.Println(getMessage("INFO", fileName, line, fmt.Sprintf(message, args...)))
	}
}

func Warning(messages ...any) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		fileName, _ := f.FileLine(pc)
		color.Yellow.Println(getMessage("WARNING", fileName, line, messages))
	}
}
func Warningf(message string, args ...any) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		fileName, _ := f.FileLine(pc)
		color.Yellow.Println(getMessage("WARNING", fileName, line, fmt.Sprintf(message, args...)))
	}
}

func Panic(messages ...any) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		fileName, _ := f.FileLine(pc)
		color.Red.Println(getMessage("PANIC", fileName, line, messages))
		panic(getMessage("PANIC", fileName, line, messages))
	}
}
func Panicf(message string, args ...any) {
	pc, _, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if ok && f.Name() != "runtime.goexit" {
		fileName, _ := f.FileLine(pc)
		color.Red.Println(getMessage("PANIC", fileName, line, fmt.Sprintf(message, args...)))
		panic(getMessage("PANIC", fileName, line, fmt.Sprintf(message, args...)))
	}
}

func getMessage(status, fileName string, line int, messages ...any) string {
	messageArray := lo.Map(messages, func(message any, _ int) string {
		return fmt.Sprintf("[%s] %s %s:%d: %v", status, time.Now().Format("2006-01-02 15:04:05"), fileName, line, message)
	})
	return strings.Join(messageArray, "\n")
}
