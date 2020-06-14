package jobs

import (
	"fmt"
	"time"

	"github.com/qinyuanmao/gomao/logger"
	"github.com/thoas/go-funk"
)

type jobStatus int

const (
	running jobStatus = iota
	stopped
)

var (
	jobChan           chan *tickerJob
	stopChan          chan string
	finishChan        chan string
	finishAllChan     chan struct{}
	finishAllChanDown chan struct{}
)

type tickerJob struct {
	jobName    string
	jobFunc    func()
	ticker     *time.Ticker
	status     jobStatus
	delayCount int
}

func init() {
	var jobsMap = make(map[string](*tickerJob))
	jobChan = make(chan *tickerJob)
	stopChan = make(chan string)
	finishChan = make(chan string)
	finishAllChan = make(chan struct{})
	finishAllChanDown = make(chan struct{})

	// 这里不能用 for-select, 会导致 map 被锁，无法更新
	go func() {
		for job := range jobChan {
			if _, has := jobsMap[job.jobName]; has {
				StopJob(job.jobName)
			}
			jobsMap[job.jobName] = job
			go jobsMap[job.jobName].startJob()
		}
	}()
	go func() {
		for jobName := range stopChan {
			if _, has := jobsMap[jobName]; has {
				jobsMap[jobName].status = stopped
				delete(jobsMap, jobName)
			} else {
				finishChan <- jobName
			}
		}
	}()
	go func() {
		<-finishAllChan
		var jobNames = funk.Map(jobsMap, func(jobName string, job *tickerJob) string {
			return jobName
		}).([]string)
		funk.ForEach(jobNames, func(jobName string) {
			StopJob(jobName)
		})
		finishAllChanDown <- struct{}{}
	}()
}

func (job *tickerJob) startJob() {
	if job == nil {
		return
	}
	defer func() {
		job.ticker.Stop()
		logger.Info(fmt.Sprintf("%s job is stopped.", job.jobName))
		if job.delayCount == 0 { // delayCount = 0 表示是手动结束的任务
			finishChan <- job.jobName
		}
	}()
	logger.Info(fmt.Sprintf("%s job is started.", job.jobName))
	var count = 0
	for {
		if job.status == running {
			<-job.ticker.C
			job.jobFunc()
			count++
			if job.delayCount != 0 && count >= job.delayCount {
				if count >= job.delayCount {
					stopChan <- job.jobName
					return
				}
			}
		} else {
			return
		}
	}
}

func StartJob(jobName string, delaytime time.Duration, delayCount int, fun func()) {
	jobChan <- &tickerJob{
		jobName:    jobName,
		jobFunc:    fun,
		ticker:     time.NewTicker(delaytime),
		delayCount: delayCount,
		status:     running,
	}
}

func StopJob(jobName string) {
	logger.Info(fmt.Sprintf("%s job is stopping.", jobName))
	stopChan <- jobName
	for jobName == <-finishChan {
		return
	}
}

func Destory() {
	finishAllChan <- struct{}{}
	<-finishAllChanDown
	close(jobChan)
	close(stopChan)
	close(finishChan)
	close(finishAllChan)
	close(finishAllChanDown)
}
