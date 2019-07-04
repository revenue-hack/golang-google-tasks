package src

import (
	"log"

	"google.golang.org/api/tasks/v1"
)

type Task struct {
	todoID, title, notes, due string
}

func NewTask(todoID, title, notes, due string) *Task {
	if todoID == "" || title == "" || notes == "" || due == "" {
		return nil
	}

	return &Task{
		todoID: todoID,
		title:  title,
		notes:  notes,
		due:    due,
	}
}

type TaskOperator interface {
	ListByTODOID(id string) []*tasks.Task
	CreateByTODOID(task *Task) *tasks.Task
	FindByTitle(id, title string) *tasks.Task
	UpdateByTODOID(task *Task) *tasks.Task
	DeleteByTitle(todoID, title string)
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

func (op *TaskOperation) CreateByTODOID(task *Task) *tasks.Task {
	t, err := op.srv.Insert(task.todoID, &tasks.Task{Title: task.title, Notes: task.notes, Due: task.due}).Do()
	if err != nil {
		log.Fatalf("Unable to create task. %v", err)
	}

	return t
}

func (op *TaskOperation) FindByTitle(id, title string) *tasks.Task {
	list := op.ListByTODOID(id)

	if len(list) > 0 {
		for _, l := range list {
			if l.Title == title {
				return l
			}
			return nil
		}
	}

	return nil
}

func (op *TaskOperation) UpdateByTODOID(task *Task) *tasks.Task {
	prevTask := op.FindByTitle(task.todoID, task.title)

	t, err := op.srv.Update(task.todoID, prevTask.Id,
		&tasks.Task{Id: prevTask.Id, Title: task.title, Due: task.due, Notes: task.notes}).Do()
	if err != nil {
		log.Fatalf("Unable to update task. %v", err)
	}

	return t
}

func (op *TaskOperation) DeleteByTitle(todoID, title string) {
	task := op.FindByTitle(todoID, title)
	if err := op.srv.Delete(todoID, task.Id).Do(); err != nil {
		log.Fatalf("Unable to delete task. %v", err)
	}
}
