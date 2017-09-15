package team

import (
	"time"
)

func (team *Team) makeReaper(finJobs chan Job, aggJobs chan []Job) {
	go team.reapJobs(finJobs, aggJobs)
}

func (team *Team) reapJobs(finJobs chan Job, aggJobs chan []Job) {
	team.Printf("R: reapJobs()")

	jobs := make([]Job, 0)

	for i := 0; i < len(team.tasks); i++ {
		job := <-finJobs
		team.Printf("R: Reap job!")

		team.average += job.elasped

		jobs = append(jobs, job)
	}

	team.average = time.Duration(int(team.average) / len(team.tasks))

	aggJobs <- jobs
}
