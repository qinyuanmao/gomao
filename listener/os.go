package listener

import (
	"os"
	"os/signal"
)

func StopApplication(action func()) {
	// 监听程序关闭
	c := make(chan os.Signal, 1)
	defer close(c)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	if sig == os.Interrupt || sig == os.Kill {
		action()
	}
}
