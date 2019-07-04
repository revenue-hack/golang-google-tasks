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
	Create(title string) *tasks.TaskList
	DeleteByTODOID(title string)
	UpdateTitleByTODOID(prevTitle, nextTitle string) *tasks.TaskList
	FindByTitle(title string) *tasks.TaskList
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

func (op *TODOOperation) FindByTitle(title string) *tasks.TaskList {
	list := op.List()

	if len(list) > 0 {
		for _, l := range list {
			if l.Title == title {
				return l
			}
		}
		return nil
	} else {
		return nil
	}
}

func (op *TODOOperation) Create(title string) *tasks.TaskList {
	tl, err := op.srv.Insert(&tasks.TaskList{Title: title}).Do()
	if err != nil {
		log.Fatalf("Unable to create todo. %v", err)
	}

	return tl
}

func (op *TODOOperation) DeleteByTODOID(title string) {
	todo := op.FindByTitle(title)

	if err := op.srv.Delete(todo.Id).Do(); err != nil {
		log.Fatalf("Unable to delete todo. %v", err)
	}
}

func (op *TODOOperation) UpdateTitleByTODOID(prevTitle, nextTitle string) *tasks.TaskList {
	prevTODO := op.FindByTitle(prevTitle)
	if prevTODO == nil {
		log.Fatal("No TODO exists")
	}

	tl, err := op.srv.Update(prevTODO.Id, &tasks.TaskList{Id: prevTODO.Id, Title: nextTitle}).Do()
	if err != nil {
		log.Fatalf("Unable to update todo. %v", err)
	}

	return tl
}
