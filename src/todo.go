package src

import (
	"golang.org/x/xerrors"
	"google.golang.org/api/tasks/v1"
)

const (
	maxCount = 10
)

type TODOOperator interface {
	List() ([]*tasks.TaskList, error)
	First() (*tasks.TaskList, error)
	Create(title string) (*tasks.TaskList, error)
	DeleteByTODOID(title string) error
	UpdateTitleByTODOID(prevTitle, nextTitle string) (*tasks.TaskList, error)
	FindByTitle(title string) (*tasks.TaskList, error)
}

type TODOOperation struct {
	wrap TODOOpWrapper
}

func NewTODOOperation(wrap TODOOpWrapper) TODOOperator {
	return &TODOOperation{wrap: wrap}
}

func (op *TODOOperation) List() ([]*tasks.TaskList, error) {
	r, err := op.wrap.List(maxCount)
	if err != nil {
		return nil, xerrors.Errorf("Unable to retrieve task lists. %v", err)
	}

	return r.Items, nil
}

func (op *TODOOperation) First() (*tasks.TaskList, error) {
	list, err := op.List()
	if err != nil {
		return nil, xerrors.Errorf("Unable to first todo, %v", err)
	}

	if len(list) > 0 {
		return list[0], nil
	}
	return nil, nil
}

func (op *TODOOperation) FindByTitle(title string) (*tasks.TaskList, error) {
	list, err := op.List()
	if err != nil {
		return nil, xerrors.Errorf("Unable to first todo, %v", err)
	}

	if len(list) > 0 {
		for _, l := range list {
			if l.Title == title {
				return l, nil
			}
		}
		return nil, nil
	}

	return nil, nil
}

func (op *TODOOperation) Create(title string) (*tasks.TaskList, error) {
	tl, err := op.wrap.Insert(&tasks.TaskList{Title: title})
	if err != nil {
		return nil, xerrors.Errorf("Unable to create todo. %v", err)
	}

	return tl, nil
}

func (op *TODOOperation) DeleteByTODOID(title string) error {
	todo, err := op.FindByTitle(title)
	if err != nil {
		return err
	}

	if err := op.wrap.Delete(todo.Id); err != nil {
		return xerrors.Errorf("Unable to delete todo. %v", err)
	}
	return nil
}

func (op *TODOOperation) UpdateTitleByTODOID(prevTitle, nextTitle string) (*tasks.TaskList, error) {
	prevTODO, err := op.FindByTitle(prevTitle)
	if err != nil {
		return nil, err
	}

	if prevTODO == nil {
		return nil, xerrors.New("No TODO exists")
	}

	tl, err := op.wrap.Update(prevTODO.Id, &tasks.TaskList{Id: prevTODO.Id, Title: nextTitle})
	if err != nil {
		return nil, xerrors.Errorf("Unable to update todo. %v", err)
	}

	return tl, nil
}
