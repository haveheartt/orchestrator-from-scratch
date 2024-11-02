package worker

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/haveheartt/orchestrator-from-scratch/task"
)

type Worker struct {
    Name        string
    Queue       queue.Queue
    Db          map[uuid.UUID]*task.Task
    TaskCount   int
}

func (w *Worker) CollectStats() {
    fmt.Println("i will collect stats")
}

func (w *Worker) RunTask() {
    fmt.Println("i will start or stop a task")    
}

func (w *Worker) StartTask() {
    fmt.Println("i will start a task")
}

func (w *Worker) StopTask() {
    fmt.Println("i will stop a task")
}


