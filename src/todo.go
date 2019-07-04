package src

import (
	"log"

	"google.golang.org/api/tasks/v1"
)

const (
	maxCount = 10
)

type TODOOperator interface {
	List() []*tasks.TaskList
	First() *tasks.TaskList
}

type TODOOperation struct {
	srv *tasks.TasklistsService
}

func NewTODOOperation(service *tasks.TasklistsService) TODOOperator {
	return &TODOOperation{srv: service}
}

func (op *TODOOperation) List() []*tasks.TaskList {
	r, err := op.srv.List().MaxResults(maxCount).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve task lists. %v", err)
	}

	return r.Items
}

func (op *TODOOperation) First() *tasks.TaskList {
	list := op.List()

	if len(list) > 0 {
		return list[0]
	} else {
		return nil
	}
}
