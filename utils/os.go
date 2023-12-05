package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func StopApplication(action func()) {
	// 监听程序关闭
	c := make(chan os.Signal, 1)
	defer close(c)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	sig := <-c
	if sig == os.Interrupt || sig == syscall.SIGTERM {
		action()
	}
}
