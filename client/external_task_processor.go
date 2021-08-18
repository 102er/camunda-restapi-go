package client

import (
	"github.com/blossom102er/camunda-restapi-go/services"
	"github.com/blossom102er/go-pool"
	"time"
)

// Reference documents：
// https://docs.camunda.org/manual/7.15/reference/rest/external-task/fetch/

const MaxLongPollTimeMinute = time.Minute * 30
const DefaultLongPollTimeMinute = time.Minute * 5

type ILogger interface {
	Debug(args ...interface{})
	Infof(template string, args ...interface{})
}

type ProcessorOptions struct {
	WorkerId                  string
	LockDuration              time.Duration
	MaxTasks                  int
	MaxParallelTaskPerHandler int
	UsePriority               *bool
	LongPollingTimeout        time.Duration
}

type ExternalTaskProcessor struct {
	client  *CamundaRestApiClient
	options *ProcessorOptions
}

func ExternalTaskRoutine(workers int, topics services.QueryFetchAndLockTopic, handleFunc gopool.HandleFunctions,
	stopSignal chan struct{}, options *ProcessorOptions, logger ILogger) {
	client := GetCamundaRestApiClient()
	//start go pool
	processor := gopool.CreateCoroutinePool(workers, handleFunc, "externalTaskRoutine")
	processor.Start()
	//The lock time of the subject task should not be less than the expected task execution time,
	//otherwise the lock failure will cause the task to be locked by other coroutines
	if topics.LockDuration <= 0 {
		topics.LockDuration = int(options.LockDuration / time.Millisecond)
	}
	//Long polling timeout time, the maximum cannot exceed 30 minutes, the default is 5 minutes
	var asyncResponseTimeout *int
	var mp int
	if options.LongPollingTimeout > MaxLongPollTimeMinute {
		mp = int(MaxLongPollTimeMinute.Milliseconds())
	} else if options.LongPollingTimeout.Nanoseconds() > 0 {
		mp = int(options.LongPollingTimeout.Nanoseconds() / int64(time.Millisecond))
	} else {
		mp = int(DefaultLongPollTimeMinute.Milliseconds())
	}
	asyncResponseTimeout = &mp
	//pull camunda external task
	retries := 0
	for {
		//Receive task stop signal
		select {
		case <-stopSignal:
			logger.Infof("topic %s receive the signal and stop the task", topics.TopicName)
			processor.Stop() //stop current topic worker
			goto end
		default:
			//Continue to perform the task by default
			logger.Infof("topic %s start fetch and lock task,lock time %d,long poll time %d", topics.TopicName, topics.LockDuration, asyncResponseTimeout)
		}
		//fetch and lock external task
		//Maintain long polling until there is a task to return,
		//will not disconnect or exceed the maximum timeout set by the long connection AsyncResponseTimeout will disconnect and enter the next cycle
		tasks, err := client.ExternalTaskService.FetchAndLock(services.QueryFetchAndLock{
			WorkerId:             options.WorkerId,
			MaxTasks:             options.MaxTasks,
			UsePriority:          options.UsePriority,
			AsyncResponseTimeout: asyncResponseTimeout,
			Topics:               &[]services.QueryFetchAndLockTopic{topics},
		})
		if err != nil {
			if retries < 60 {
				retries += 1
			}
			logger.Infof("topic %s failed pull: %v, sleeping: %d seconds", topics.TopicName, err, retries)
			time.Sleep(time.Duration(retries) * time.Second)
			continue
		}
		retries = 0
		logger.Infof("start handle topic %s,task count=%d", topics.TopicName, len(tasks))
		for _, task := range tasks {
			processor.PublishTask(task)
		}
	}
end:
	logger.Infof("exit external task client,topic:%s", topics.TopicName)
}
