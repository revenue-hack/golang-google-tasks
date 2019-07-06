package src_test

import (
	"errors"
	"reflect"
	"testing"

	"golang.org/x/xerrors"

	"github.com/revenue-hack/golang-google-tasks/src"

	"github.com/golang/mock/gomock"
	"github.com/revenue-hack/golang-google-tasks/src/wrappermock"
	"google.golang.org/api/tasks/v1"
)

func TestTaskOperation_CreateByTODOID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wrapper := wrappermock.NewMockTaskOpWrapper(ctrl)

	todoID := "id"
	title := "title"
	notes := "description"
	due := "2019-01-01"
	in := src.NewTask(todoID, title, notes, due)

	t.Run("of normal", func(t *testing.T) {
		expected := &tasks.Task{Id: "111", Title: "test"}
		task := &tasks.Task{Title: title, Due: due, Notes: notes}
		wrapper.EXPECT().Insert(todoID, task).Return(expected, nil)

		op := src.NewTaskOperation(wrapper)
		result, err := op.CreateByTODOID(in)

		if err != nil {
			t.Errorf("expected err is nil, result err is %v", err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("match error, result is %v, expected id %v", result, expected)
		}
	})

	t.Run("of abnormal", func(t *testing.T) {
		task := &tasks.Task{Title: title, Due: due, Notes: notes}
		wrapper.EXPECT().Insert(todoID, task).Return(nil, errors.New("abnormal"))

		op := src.NewTaskOperation(wrapper)
		if _, err := op.CreateByTODOID(in); err == nil {
			t.Errorf("expected err is not nil")
		}
	})
}

func TestTaskOperation_DeleteByTitle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskID := "taskid"
	todoID := "todoid"
	title := "title"
	tasks := &tasks.Tasks{Items: []*tasks.Task{
		{
			Id:    taskID,
			Title: title,
		},
	}}

	wrapper := wrappermock.NewMockTaskOpWrapper(ctrl)

	t.Run("of normal", func(t *testing.T) {
		wrapper.EXPECT().Delete(todoID, taskID).Return(nil)
		wrapper.EXPECT().List(todoID).Return(tasks, nil)

		op := src.NewTaskOperation(wrapper)
		if err := op.DeleteByTitle(todoID, title); err != nil {
			t.Errorf("expected err is nil, err is %v", err)
		}
	})

	t.Run("of abnormal, find by title error", func(t *testing.T) {
		wrapper.EXPECT().List(todoID).Return(nil, xerrors.New("err"))
		op := src.NewTaskOperation(wrapper)
		if err := op.DeleteByTitle(todoID, title); err == nil {
			t.Error("expected err is not nil,")
		}
	})

	t.Run("of abnormal, delete error", func(t *testing.T) {
		wrapper.EXPECT().Delete(todoID, taskID).Return(xerrors.New("err"))
		wrapper.EXPECT().List(todoID).Return(tasks, nil)

		op := src.NewTaskOperation(wrapper)
		if err := op.DeleteByTitle(todoID, title); err == nil {
			t.Error("expected err is not nil")
		}
	})
}
func TestTODOOperation_UpdateTitleByTODOID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wrapper := wrappermock.NewMockTaskOpWrapper(ctrl)

	todoID := "taskid"
	title := "title"
	notes := "description"
	due := "2019-01-01"
	taskID := "taskid"
	task := &tasks.Task{Id: taskID, Title: title, Due: due, Notes: notes}
	taskList := &tasks.Tasks{Items: []*tasks.Task{
		{
			Id:    taskID,
			Title: title,
		},
	}}

	t.Run("of normal", func(t *testing.T) {
		expectedTask := &tasks.Task{Id: "id"}
		wrapper.EXPECT().List(todoID).Return(taskList, nil)
		wrapper.EXPECT().Update(todoID, taskID, task).Return(expectedTask, nil)

		op := src.NewTaskOperation(wrapper)
		result, err := op.UpdateByTODOID(src.NewTask(todoID, title, notes, due))
		if err != nil {
			t.Errorf("expected err is nil, err is %v", err)
		}
		if !reflect.DeepEqual(result, expectedTask) {
			t.Errorf("expected is %v, result is %v", expectedTask, result)
		}
	})

	t.Run("of abnormal, find by title error", func(t *testing.T) {
		wrapper.EXPECT().List(todoID).Return(nil, xerrors.New("err"))

		op := src.NewTaskOperation(wrapper)
		_, err := op.UpdateByTODOID(src.NewTask(todoID, title, notes, due))
		if err == nil {
			t.Error("expected err is not nil")
		}
	})

	t.Run("of abnormal, update error", func(t *testing.T) {
		wrapper.EXPECT().List(todoID).Return(taskList, nil)
		wrapper.EXPECT().Update(todoID, taskID, task).Return(nil, xerrors.New("err"))

		op := src.NewTaskOperation(wrapper)
		_, err := op.UpdateByTODOID(src.NewTask(todoID, title, notes, due))
		if err == nil {
			t.Error("expected err is not nil")
		}
	})
}

func TestTaskOperation_ListByTODOID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wrapper := wrappermock.NewMockTaskOpWrapper(ctrl)

	todoID := "taskid"
	title := "title"
	taskID := "taskid"
	taskList := &tasks.Tasks{Items: []*tasks.Task{
		{
			Id:    taskID,
			Title: title,
		},
	}}

	t.Run("of normal", func(t *testing.T) {
		wrapper.EXPECT().List(todoID).Return(taskList, nil)

		op := src.NewTaskOperation(wrapper)
		result, err := op.ListByTODOID(todoID)
		if err != nil {
			t.Errorf("expected err is nil, err is %v", err)
		}
		if !reflect.DeepEqual(result, taskList.Items) {
			t.Errorf("unmatch err. expected is %v, result is %v", taskList, result)
		}
	})

	t.Run("of abnormal", func(t *testing.T) {
		wrapper.EXPECT().List(todoID).Return(nil, xerrors.New("abnormal"))

		op := src.NewTaskOperation(wrapper)
		if _, err := op.ListByTODOID(todoID); err == nil {
			t.Errorf("expected err is not nil")
		}
	})
}
