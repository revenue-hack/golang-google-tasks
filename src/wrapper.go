// mockgen
package src

import tasks "google.golang.org/api/tasks/v1"

type TODOOpWrap struct {
	srv *tasks.TasklistsService
}

type TODOOpWrapper interface {
	List(maxCount int64) *tasks.TasklistsListCall
	Insert(tasklist *tasks.TaskList) *tasks.TasklistsInsertCall
	Update(tasklistid string, tasklist *tasks.TaskList) *tasks.TasklistsUpdateCall
	Delete(tasklistid string) *tasks.TasklistsDeleteCall
}

func NewTODOOpWrap(srv *tasks.TasklistsService) TODOOpWrapper {
	return &TODOOpWrap{srv}
}

func (op *TODOOpWrap) List(maxCount int64) *tasks.TasklistsListCall {
	return op.srv.List().MaxResults(maxCount)
}

func (op *TODOOpWrap) Insert(tasklist *tasks.TaskList) *tasks.TasklistsInsertCall {
	return op.srv.Insert(tasklist)
}

func (op *TODOOpWrap) Update(tasklistid string, tasklist *tasks.TaskList) *tasks.TasklistsUpdateCall {
	return op.srv.Update(tasklistid, tasklist)
}

func (op *TODOOpWrap) Delete(tasklistid string) *tasks.TasklistsDeleteCall {
	return op.srv.Delete(tasklistid)
}

type TaskOpWrap struct {
	srv *tasks.TasksService
}

type TaskOpWrapper interface {
	List(tasklistid string) *tasks.TasksListCall
	Insert(tasklistid string, task *tasks.Task) *tasks.TasksInsertCall
	Update(tasklistid string, taskid string, task *tasks.Task) *tasks.TasksUpdateCall
	Delete(tasklistid string, taskid string) *tasks.TasksDeleteCall
}

func NewTaskOpWrap(srv *tasks.TasksService) TaskOpWrapper {
	return &TaskOpWrap{srv}
}

func (op *TaskOpWrap) List(tasklistid string) *tasks.TasksListCall {
	return op.srv.List(tasklistid)
}

func (op *TaskOpWrap) Insert(tasklistid string, task *tasks.Task) *tasks.TasksInsertCall {
	return op.srv.Insert(tasklistid, task)
}

func (op *TaskOpWrap) Update(tasklistid string, taskid string, task *tasks.Task) *tasks.TasksUpdateCall {
	return op.srv.Update(tasklistid, taskid, task)
}

func (op *TaskOpWrap) Delete(tasklistid string, taskid string) *tasks.TasksDeleteCall {
	return op.srv.Delete(tasklistid, taskid)
}
