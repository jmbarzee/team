package main

/*	----------------------== Ussage of team ==------------------------
 *	There are only four nessecarry bits of code to run the team pattern
 *	You can find all points here.
 *	Copy, paste, and modify for quick implementation of go concurency patterns
 */

import (
	"fmt"
	"github.com/jmbarzee/team"
	"log"
	"os"
	"runtime"
	"time"
)

type myTask struct {
	doTask func()
}

func (t *myTask) Run() {
	t.doTask()
}

func main() {
	coreNum := 8
	runtime.GOMAXPROCS(coreNum) // Set number of cores
	fmt.Printf("Running on %v cores", coreNum)

	tasks := buildTasks()
	team.NewTeam(100000, 100000).Run(tasks)
}

func buildTasks() []team.Task {
	logger := log.New(os.Stdout, "Main", log.LstdFlags) // we need a thread safe PrintF
	tasks := make([]team.Task, 0)
	numTasks := 10000000
	for i := 0; i < numTasks; i++ {
		taskNum := i
		tasks = append(tasks, &myTask{
			doTask: func() {
				// simulate waiting for something we are not responsible for.
				// i.e. a network request
				time.Sleep(time.Second)
				// indicate which task number we were
				logger.Printf("%v", taskNum)
			},
		})
	}
	return tasks
}
