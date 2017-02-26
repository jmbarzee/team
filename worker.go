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

/*
func (fs *fSetter) req(url string) (int, error) {
	//fs.DPrintf("        URL: %v", url)
	var requestType string
	var jsonBody []byte
	var err error
	if fs.setting == -1 {
		requestType = "DELETE"
		jsonBody = []byte("{}")
	} else {
		requestType = "PUT"
		jsonBody, err = json.Marshal(struct {
			Dial   int    `json:"dial"`
			Author string `json:"author"`
		}{
			Dial:   fs.setting,
			Author: fs.author,
		})

	}
	req, err := http.NewRequest(requestType, url, bytes.NewReader(jsonBody))
	if err != nil {
		fmt.Errorf("%#v", err)
		return 0, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	//fs.DPrintf("        resp: %v", resp)
	defer resp.Body.Close()
	return resp.StatusCode, err
}
*/
