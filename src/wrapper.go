// mockgen
package src

import (
	"google.golang.org/api/googleapi"
	tasks "google.golang.org/api/tasks/v1"
)

type TODOOpWrap struct {
	srv *tasks.TasklistsService
}

type TODOOpWrapper interface {
	List(maxCount int64, opts ...googleapi.CallOption) (*tasks.TaskLists, error)
	Insert(tasklist *tasks.TaskList, opts ...googleapi.CallOption) (*tasks.TaskList, error)
	Update(tasklistid string, tasklist *tasks.TaskList, opts ...googleapi.CallOption) (*tasks.TaskList, error)
	Delete(tasklistid string, opts ...googleapi.CallOption) error
}

func NewTODOOOpWrap(srv *tasks.TasklistsService) TODOOpWrapper {
	return &TODOOpWrap{srv}
}

func (op *TODOOpWrap) List(maxCount int64, opts ...googleapi.CallOption) (*tasks.TaskLists, error) {
	return op.srv.List().MaxResults(maxCount).Do(opts...)
}

func (op *TODOOpWrap) Insert(tasklist *tasks.TaskList, opts ...googleapi.CallOption) (*tasks.TaskList, error) {
	return op.srv.Insert(tasklist).Do(opts...)
}

func (op *TODOOpWrap) Update(tasklistid string, tasklist *tasks.TaskList, opts ...googleapi.CallOption) (*tasks.TaskList, error) {
	return op.srv.Update(tasklistid, tasklist).Do(opts...)
}

func (op *TODOOpWrap) Delete(tasklistid string, opts ...googleapi.CallOption) error {
	return op.srv.Delete(tasklistid).Do(opts...)
}

type TaskOpWrap struct {
	srv *tasks.TasksService
}

type TaskOpWrapper interface {
	List(tasklistid string, opts ...googleapi.CallOption) (*tasks.Tasks, error)
	Insert(tasklistid string, task *tasks.Task, opts ...googleapi.CallOption) (*tasks.Task, error)
	Update(tasklistid string, taskid string, task *tasks.Task, opts ...googleapi.CallOption) (*tasks.Task, error)
	Delete(tasklistid string, taskid string, opts ...googleapi.CallOption) error
}

func NewTaskOpWrap(srv *tasks.TasksService) TaskOpWrapper {
	return &TaskOpWrap{srv}
}

func (op *TaskOpWrap) List(tasklistid string, opts ...googleapi.CallOption) (*tasks.Tasks, error) {
	return op.srv.List(tasklistid).Do(opts...)
}

func (op *TaskOpWrap) Insert(tasklistid string, task *tasks.Task, opts ...googleapi.CallOption) (*tasks.Task, error) {
	return op.srv.Insert(tasklistid, task).Do(opts...)
}

func (op *TaskOpWrap) Update(tasklistid string,
	taskid string,
	task *tasks.Task,
	opts ...googleapi.CallOption) (*tasks.Task, error) {
	return op.srv.Update(tasklistid, taskid, task).Do(opts...)
}

func (op *TaskOpWrap) Delete(tasklistid string, taskid string, opts ...googleapi.CallOption) error {
	return op.srv.Delete(tasklistid, taskid).Do(opts...)
}
