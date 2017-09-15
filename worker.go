package team

func (team *Team) makeWorkers(newJobs chan Job, finJobs chan Job, workerCount int) {
	for i := 0; i < workerCount; i++ {
		go team.workJobs(newJobs, finJobs)
	}
}

func (team *Team) workJobs(newJobs chan Job, finJobs chan Job) {
	team.Printf("W: workJobs()")
	for job := range newJobs {
		team.Printf("W:  Got job!")
		job.Work()
		finJobs <- job
	}
}
