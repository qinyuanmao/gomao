package main

import (
	"fmt"
	"time"

	"e.coding.net/tssoft/repository/gomao/listener"
	"e.coding.net/tssoft/repository/gomao/logger"

	"e.coding.net/tssoft/repository/gomao/jobs"
	_ "e.coding.net/tssoft/repository/gomao/jobs"
)

func main() {
	// 定时任务，永远执行
	jobs.StartJob("test1", 1*time.Second, 0, func() {
		fmt.Println("test1 is running")
	})

	time.Sleep(5 * time.Second)
	// 重复 job 执行，先关闭原来的，再执行新的
	jobs.StartJob("test1", 1*time.Second, 0, func() {
		fmt.Println("test1—new is running")
	})

	// 定时任务，只执行 5 次，非永远任务不需要关闭，执行 5 次之后会自动停止
	jobs.StartJob("test2", 1*time.Second, 5, func() {
		fmt.Println("test2 is running")
	})

	// 定时任务，永远执行
	jobs.StartJob("test3", 1*time.Second, 0, func() {
		fmt.Println("test3 is running")
	})

	// 定时任务，永远执行
	jobs.StartJob("test4", 10*time.Second, 0, func() {
		fmt.Println("test4 is running")
	})

	listener.StopApplication(func() {
		jobs.StopJob("test1")
		jobs.StopJob("test2")
		jobs.StopJob("test5") // 结束不存在的 job
		jobs.Destory()
		logger.Info("stop jobs")
	})
}
