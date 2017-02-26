package main

import (
	"time"
)

func (team *Team) makeSender(newJobs chan Job, maxJobsPerSec int) {
	go team.sendJobs(newJobs, maxJobsPerSec)
}

func (team *Team) sendJobs(newJobs chan Job, maxJobsPerSec int) {
	team.Printf("S: sendJobs()")

	NSecsPerTic := time.Duration(int(time.Second) / maxJobsPerSec)
	ticker := time.NewTicker(NSecsPerTic)

	for _, task := range team.tasks {
		<-ticker.C                 // wait for a tick
		newJobs <- Job{task: task} // put a job in the channel
		team.Printf("S: Sent job!")
	}

	close(newJobs) // shut down channel
	ticker.Stop()  // shut off ticker
}
