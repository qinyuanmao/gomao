package logger

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/samber/lo"
)

type Logger struct {
	level      int
	startLevel int
	debug      bool
}

func New(level int, debug bool) *Logger {
	return newLevel(level, 2, debug) // 2 是 getCaller() 的上一层级
}

func newLevel(level int, startLevel int, debug bool) *Logger {
	return &Logger{level: level, startLevel: startLevel, debug: debug}
}

func (logger *Logger) getCaller() string {
	var array []string
	for i := logger.startLevel; i < logger.level+logger.startLevel; i++ {
		pc, _, line, ok := runtime.Caller(i)
		f := runtime.FuncForPC(pc)
		if ok && f.Name() != "runtime.goexit" {
			fileName, _ := f.FileLine(pc)
			if i != 1 {
				array = append(array, fmt.Sprintf("\t%s:%d", fileName, line))
			} else {
				array = append(array, fmt.Sprintf("%s:%d", fileName, line))
			}
		} else {
			break
		}
	}
	return strings.Join(array, "\n")
}

func (logger *Logger) Error(messages ...any) {
	color.Red.Println(getMessage("ERROR", logger.getCaller(), messages))
}
func (logger *Logger) Errorf(message string, args ...any) {
	color.Red.Println(getMessage("ERROR", logger.getCaller(), fmt.Sprintf(message, args...)))
}

func (logger *Logger) Debug(messages ...any) {
	color.Green.Println(getMessage("DEBUG", logger.getCaller(), messages))
}
func (logger *Logger) Debugf(message string, args ...any) {
	color.Green.Println(getMessage("DEBUG", logger.getCaller(), fmt.Sprintf(message, args...)))
}

func (logger *Logger) Info(messages ...any) {
	color.White.Println(getMessage("INFO", logger.getCaller(), messages))
}
func (logger *Logger) Infof(message string, args ...any) {
	color.White.Println(getMessage("INFO", logger.getCaller(), fmt.Sprintf(message, args...)))
}

func (logger *Logger) Warning(messages ...any) {
	color.Yellow.Println(getMessage("WARNING", logger.getCaller(), messages))
}
func (logger *Logger) Warningf(message string, args ...any) {
	color.Yellow.Println(getMessage("WARNING", logger.getCaller(), fmt.Sprintf(message, args...)))
}

func (logger *Logger) Panic(messages ...any) {
	color.Red.Println(getMessage("PANIC", logger.getCaller(), messages))
	panic(getMessage("PANIC", logger.getCaller(), messages))
}
func (logger *Logger) Panicf(message string, args ...any) {
	color.Red.Println(getMessage("PANIC", logger.getCaller(), fmt.Sprintf(message, args...)))
	panic(getMessage("PANIC", logger.getCaller(), fmt.Sprintf(message, args...)))
}

func (logger *Logger) getMessage(status, fileName string, line int, messages ...any) string {
	messageArray := lo.Map(messages, func(message any, _ int) string {
		return fmt.Sprintf("[%s] %s %s:%d: %v", status, time.Now().Format("2006-01-02 15:04:05"), fileName, line, message)
	})
	return strings.Join(messageArray, "\n")
}
