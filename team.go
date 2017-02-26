package main

import (
	"log"
	"os"
	"time"
)

type Team struct {
	logger *log.Logger

	tasks []Task

	start   time.Time
	end     time.Time
	elasped time.Duration
	average time.Duration

	maxParellelJobs int
	maxJobsPerSec   int
}

const queueSizeMult = 2

func NewTeam(maxParellelJobs int, maxJobsPerSec int) *Team {
	team := Team{
		logger: log.New(os.Stdout, "Team", log.LstdFlags),

		maxParellelJobs: maxParellelJobs,
		maxJobsPerSec:   maxJobsPerSec,
	}
	return &team
}

func (team *Team) Run(tasks []Task) {
	team.start = time.Now() // record start time
	team.tasks = tasks

	newJobs := make(chan Job, team.maxParellelJobs*queueSizeMult)
	finJobs := make(chan Job, team.maxParellelJobs*queueSizeMult)
	aggJobs := make(chan []Job)

	team.makeSender(newJobs, team.maxJobsPerSec)
	team.makeWorkers(newJobs, finJobs, team.maxParellelJobs)
	team.makeReaper(finJobs, aggJobs)

	<-aggJobs // wait on job reaper

	team.end = time.Now() // record start time
	team.elasped = team.end.Sub(team.start)

	team.logger.Printf("Jobs: %v", len(team.tasks))
	//team.logger.Printf("Time: %v -> %v", team.start., team.end)
	team.logger.Printf("Elasped: %v", team.elasped)
	team.logger.Printf("Average: %v", team.average)
	team.logger.Printf("jobs/s: %v", float64(len(team.tasks))/float64(team.elasped.Seconds()))
}

func (team *Team) Printf(format string, a ...interface{}) {
	team.logger.Printf(format, a...)
}
