# go-quartz
[![Build Status](https://travis-ci.org/reugn/go-quartz.svg?branch=master)](https://travis-ci.org/reugn/go-quartz)
[![GoDoc](https://godoc.org/github.com/reugn/go-quartz/quartz?status.svg)](https://godoc.org/github.com/reugn/go-quartz/quartz)
[![Go Report Card](https://goreportcard.com/badge/github.com/reugn/go-quartz)](https://goreportcard.com/report/github.com/reugn/go-quartz)
[![codecov](https://codecov.io/gh/reugn/go-quartz/branch/master/graph/badge.svg)](https://codecov.io/gh/reugn/go-quartz)

A minimalistic and zero-dependency scheduling library for Go.

## About
Inspired by the [Quartz](https://github.com/quartz-scheduler/quartz) Java scheduler.

### Library building blocks
Job interface. Any type that implements it can be scheduled.
```go
type Job interface {
	Execute()
	Description() string
	Key() int
}
```
Implemented Jobs
- ShellJob
- CurlJob

Scheduler interface
```go
type Scheduler interface {
	// start the scheduler
	Start()
	// whether the scheduler has been started
	IsStarted() bool
	// schedule the job with the specified trigger
	ScheduleJob(job Job, trigger Trigger) error
	// get all scheduled jobs keys
	GetJobKeys() []int
	// get the scheduled job metadata
	GetScheduledJob(key int) (*ScheduledJob, error)
	// remove the job from the execution queue
	DeleteJob(key int) error
	// clear all scheduled jobs
	Clear()
	// shutdown the scheduler
	Stop()
}
```
Implemented Schedulers
- StdScheduler

Trigger interface
```go
type Trigger interface {
	NextFireTime(prev int64) (int64, error)
	Description() string
}
```
Implemented Triggers
- CronTrigger
- SimpleTrigger
- RunOnceTrigger

## Examples
```go
sched := quartz.NewStdScheduler()
sched.Start()
cronTrigger, _ := quartz.NewCronTrigger("1/5 * * * * *")
shellJob := quartz.NewShellJob("ls -la")
curlJob, _ := quartz.NewCurlJob(http.MethodGet, "http://worldclockapi.com/api/json/est/now", "", nil)
sched.ScheduleJob(shellJob, cronTrigger)
sched.ScheduleJob(curlJob, quartz.NewSimpleTrigger(time.Second*7))
sched.Stop()
```
More code samples can be found in the examples directory.

## License
Licensed under the MIT License.
