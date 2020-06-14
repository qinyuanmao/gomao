package jobs

import (
	"fmt"
	"time"

	"github.com/qinyuanmao/gomao/logger"
)

type jobStatus int

var (
	jobChan  chan *tickerJob
	stopChan chan string
)

const (
	jobRunning jobStatus = iota
	jobStopped
)

type tickerJob struct {
	jobName    string
	jobFunc    func()
	status     jobStatus
	ticker     *time.Ticker
	delayCount int
}

func init() {
	var jobsMap = make(map[string]*tickerJob)
	jobChan = make(chan *tickerJob)
	stopChan = make(chan string)
	go func() {
		for {
			select {
			case job := <-jobChan:
				if _, has := jobsMap[job.jobName]; has {
					jobsMap[job.jobName].status = jobStopped
					delete(jobsMap, job.jobName)
				}
				jobsMap[job.jobName] = job
				go startJob(jobsMap[job.jobName])
			case jobName := <-stopChan:
				jobsMap[jobName].status = jobStopped
				delete(jobsMap, jobName)
			}
		}
	}()
}

func startJob(job *tickerJob) {
	defer job.ticker.Stop()
	defer logger.Info(fmt.Sprintf("%s job is stopped.", job.jobName))
	logger.Info(fmt.Sprintf("%s job is started.", job.jobName))
	var count = 0
	for {
		if job.status == jobRunning {
			<-job.ticker.C
			job.jobFunc()
		} else {
			return
		}
		count++
		if job.delayCount != 0 && count >= job.delayCount {
			if count >= job.delayCount {
				stopChan <- job.jobName
				return
			}
		}
	}
}

func StartJob(jobName string, delaytime time.Duration, delayCount int, fun func()) {
	jobChan <- &tickerJob{
		jobName:    jobName,
		jobFunc:    fun,
		status:     jobRunning,
		ticker:     time.NewTicker(delaytime),
		delayCount: delayCount,
	}
}

func StopJob(key string) {
	stopChan <- key
}
