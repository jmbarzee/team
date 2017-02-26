package main

import (
	"time"
)

type Task interface {
	Run() // Work it
}

type Job struct {
	task    Task
	elasped time.Duration
}

func (j *Job) Work() {
	start := time.Now()
	j.task.Run()
	j.elasped = time.Since(start)
}
