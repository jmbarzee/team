package main

func (team *Team) makeWorkers(newJobs chan Job, finJobs chan Job, workerCount int) {
	for i := 0; i < workerCount; i++ {
		go team.workerJob(newJobs, finJobs)
	}
}

func (team *Team) workerJob(newJobs chan Job, finJobs chan Job) {
	team.Printf("W: workerJob()")
	for job := range newJobs {
		team.Printf("W:  Got job!")
		job.Work()
		finJobs <- job
	}
}
