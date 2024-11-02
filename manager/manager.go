package manager

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/haveheartt/orchestrator-from-scratch/task"
)

type Manager struct {
    Pending         queue.Queue
    TaskDb          map[string][]*task.Task
    EventDb         map[string][]*task.TaskEvent
    Workers         []string
    WorkerTaskMap   map[string]uuid.UUID
    TaskWorkerMap   map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
    fmt.Println("i will select an appropriate worker")
}

func (m *Manager) UpdateTasks() {
    fmt.Println("i will update tasks")
}

func (m *Manager) SendWork() {
    fmt.Println("i will send work to workers")
}


