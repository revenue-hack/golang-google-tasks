package src

import (
	"golang.org/x/xerrors"

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
	ListByTODOID(id string) ([]*tasks.Task, error)
	CreateByTODOID(task *Task) (*tasks.Task, error)
	UpdateByTODOID(task *Task) (*tasks.Task, error)
	DeleteByTitle(todoID, title string) error
}

type TaskOperation struct {
	wrap TaskOpWrapper
}

func NewTaskOperation(op TaskOpWrapper) TaskOperator {
	return &TaskOperation{wrap: op}
}

func (op *TaskOperation) ListByTODOID(id string) ([]*tasks.Task, error) {
	task, err := op.wrap.List(id)
	if err != nil {
		return nil, xerrors.Errorf("Unable to retrieve task lists. %v", err)
	}

	return task.Items, nil
}

func (op *TaskOperation) CreateByTODOID(task *Task) (*tasks.Task, error) {
	t, err := op.wrap.Insert(task.todoID, &tasks.Task{Title: task.title, Notes: task.notes, Due: task.due})
	if err != nil {
		return nil, xerrors.Errorf("Unable to create task. %v", err)
	}

	return t, nil
}

func (op *TaskOperation) findByTitle(id, title string) (*tasks.Task, error) {
	list, err := op.ListByTODOID(id)
	if err != nil {
		return nil, xerrors.Errorf("Unable to find task, %v", err)
	}

	if len(list) > 0 {
		for _, l := range list {
			if l.Title == title {
				return l, nil
			}
			return nil, nil
		}
	}

	return nil, nil
}

func (op *TaskOperation) UpdateByTODOID(task *Task) (*tasks.Task, error) {
	prevTask, err := op.findByTitle(task.todoID, task.title)
	if err != nil {
		return nil, err
	}

	t, err := op.wrap.Update(task.todoID, prevTask.Id,
		&tasks.Task{Id: prevTask.Id, Title: task.title, Due: task.due, Notes: task.notes})
	if err != nil {
		return nil, xerrors.Errorf("Unable to update task. %v", err)
	}

	return t, nil
}

func (op *TaskOperation) DeleteByTitle(todoID, title string) error {
	task, err := op.findByTitle(todoID, title)
	if err != nil {
		return err
	}

	if err := op.wrap.Delete(todoID, task.Id); err != nil {
		return xerrors.Errorf("Unable to delete task. %v", err)
	}
	return nil
}
