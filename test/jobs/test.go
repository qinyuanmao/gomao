package main

import (
	"fmt"
	"time"

	"github.com/qinyuanmao/gomao/listener"
	"github.com/qinyuanmao/gomao/logger"

	"github.com/qinyuanmao/gomao/jobs"
)

func main() {
	// 定时任务，永远执行
	jobs.StartJob("test1", 1*time.Second, 0, func() {
		fmt.Println("test1 is running")
	})
	time.Sleep(10 * time.Second)
	jobs.StopJob("test1")

	// 定时任务，只执行 5 次，非永远任务不需要关闭，执行 5 次之后会自动停止
	jobs.StartJob("test2", 1*time.Second, 5, func() {
		fmt.Println("test2 is running")
	})

	listener.StopApplication(func() {
		logger.Info("stop jobs")
	})
}
