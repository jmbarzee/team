package main

import (
	"log"
	"os"
	"runtime"
)

var words = [...]string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
}

type myTask struct {
	doTask func()
}

func (t *myTask) Run() {
	t.doTask()
}

func main() {
	logger := log.New(os.Stdout, "Main", log.LstdFlags) // we need a thread safe PrintF

	coreNum := 1
	runtime.GOMAXPROCS(coreNum) // Set number of cores
	logger.Printf("Running on %v cores", coreNum)

	tasks := make([]Task, 0)
	for _, word := range words {
		tempWord := word // to avoid scope problems
		tasks = append(tasks, &myTask{
			doTask: func() {
				logger.Printf("%v", tempWord)
			},
		})
	}
	NewTeam(2, 1).Run(tasks)
}
