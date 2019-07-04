package src

import (
	"log"

	"google.golang.org/api/tasks/v1"
)

type TaskOperator interface {
	ListByTODOID(id string) []*tasks.Task
}

type TaskOperation struct {
	srv *tasks.TasksService
}

func NewTaskOperation(srv *tasks.TasksService) TaskOperator {
	return &TaskOperation{srv: srv}
}

func (op *TaskOperation) ListByTODOID(id string) []*tasks.Task {
	task, err := op.srv.List(id).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve task lists. %v", err)
	}

	return task.Items
}
