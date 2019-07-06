package src_test

import (
	"reflect"
	"testing"

	"golang.org/x/xerrors"

	"github.com/golang/mock/gomock"
	"github.com/revenue-hack/golang-google-tasks/src"
	"github.com/revenue-hack/golang-google-tasks/src/wrappermock"
	"google.golang.org/api/tasks/v1"
)

func TestTODOOperation_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	title := "title"
	taskList := &tasks.TaskList{Title: title}

	wrapper := wrappermock.NewMockTODOOpWrapper(ctrl)

	t.Run("of normal", func(t *testing.T) {
		wrapper.EXPECT().Insert(taskList).Return(taskList, nil)

		op := src.NewTODOOperation(wrapper)
		result, err := op.Create(title)
		if err != nil {
			t.Errorf("expected err is nil, err is %v", err)
		}
		if !reflect.DeepEqual(result, taskList) {
			t.Errorf("unmatch err, result is %v, expected is %v", result, taskList)
		}
	})

	t.Run("of abnormal", func(t *testing.T) {
		wrapper.EXPECT().Insert(taskList).Return(nil, xerrors.New("err"))

		op := src.NewTODOOperation(wrapper)

		if _, err := op.Create(title); err == nil {
			t.Error("expected err is not nil")
		}
	})
}

func TestTODOOperation_First(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	title := "title"
	taskList := &tasks.TaskList{Title: title}
	taskLists := &tasks.TaskLists{Items: []*tasks.TaskList{taskList}}

	wrapper := wrappermock.NewMockTODOOpWrapper(ctrl)

	t.Run("of normal", func(t *testing.T) {
		wrapper.EXPECT().List(src.MaxCount).Return(taskLists, nil)

		op := src.NewTODOOperation(wrapper)
		result, err := op.First()
		if err != nil {
			t.Errorf("expected err is nil, err is %v", err)
		}
		if !reflect.DeepEqual(result, taskList) {
			t.Errorf("unmatch err, result is %v, expected is %v", result, taskList)
		}
	})

	t.Run("of normal", func(t *testing.T) {
		wrapper.EXPECT().List(src.MaxCount).Return(nil, xerrors.New("err"))

		op := src.NewTODOOperation(wrapper)

		if _, err := op.First(); err == nil {
			t.Error("expected err is not nil")
		}
	})
}

func TestTODOOperation_DeleteByTODOID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	title := "title"
	id := "id"
	taskList := &tasks.TaskList{Title: title, Id: id}
	taskLists := &tasks.TaskLists{Items: []*tasks.TaskList{taskList}}

	wrapper := wrappermock.NewMockTODOOpWrapper(ctrl)

	t.Run("of normal", func(t *testing.T) {
		wrapper.EXPECT().List(src.MaxCount).Return(taskLists, nil)
		wrapper.EXPECT().Delete(id).Return(nil)

		op := src.NewTODOOperation(wrapper)

		if err := op.DeleteByTitle(title); err != nil {
			t.Errorf("expected err is nil, err is %v", err)
		}
	})

	t.Run("of abnormal. find by title error", func(t *testing.T) {
		wrapper.EXPECT().List(src.MaxCount).Return(nil, xerrors.New("err"))

		op := src.NewTODOOperation(wrapper)

		if err := op.DeleteByTitle(title); err == nil {
			t.Error("expected err is not nil")
		}
	})

	t.Run("of abnormal. delete error", func(t *testing.T) {
		wrapper.EXPECT().List(src.MaxCount).Return(taskLists, nil)
		wrapper.EXPECT().Delete(id).Return(xerrors.New("err"))

		op := src.NewTODOOperation(wrapper)

		if err := op.DeleteByTitle(title); err == nil {
			t.Error("expected err is not nil")
		}
	})
}

func TestTODOOperation_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	title := "title"
	id := "id"
	taskList := &tasks.TaskList{Title: title, Id: id}
	taskLists := &tasks.TaskLists{Items: []*tasks.TaskList{taskList}}

	wrapper := wrappermock.NewMockTODOOpWrapper(ctrl)

	t.Run("of normal", func(t *testing.T) {
		wrapper.EXPECT().List(src.MaxCount).Return(taskLists, nil)

		op := src.NewTODOOperation(wrapper)
		result, err := op.List()
		if err != nil {
			t.Errorf("expected err is nil, err is %v", err)
		}
		if !reflect.DeepEqual(result, taskLists.Items) {
			t.Errorf("unmatch err, result is %v, expected is %v", result, taskList)
		}
	})

	t.Run("of abnormal", func(t *testing.T) {
		wrapper.EXPECT().List(src.MaxCount).Return(nil, xerrors.New("err"))

		op := src.NewTODOOperation(wrapper)

		if _, err := op.List(); err == nil {
			t.Error("expected err is not nil")
		}
	})
}

func TestTaskOperation_UpdateByTODOID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	title := "title"
	id := "id"
	taskList := &tasks.TaskList{Title: title, Id: id}
	taskLists := &tasks.TaskLists{Items: []*tasks.TaskList{taskList}}

	wrapper := wrappermock.NewMockTODOOpWrapper(ctrl)

	t.Run("of normal", func(t *testing.T) {
		wrapper.EXPECT().List(src.MaxCount).Return(taskLists, nil)
		wrapper.EXPECT().Update(id, taskList).Return(taskList, nil)

		op := src.NewTODOOperation(wrapper)
		result, err := op.UpdateTitleByTODOID(title, title)
		if err != nil {
			t.Errorf("expected err is nil, err is %v", err)
		}
		if !reflect.DeepEqual(result, taskList) {
			t.Errorf("unmatch error, result is %v, expected is %v", result, taskList)
		}
	})

	t.Run("of abnormal. find by title error", func(t *testing.T) {
		wrapper.EXPECT().List(src.MaxCount).Return(nil, xerrors.New("err"))

		op := src.NewTODOOperation(wrapper)

		if _, err := op.UpdateTitleByTODOID(title, title); err == nil {
			t.Error("expected err is not nil")
		}
	})

	t.Run("of abnormal. update error", func(t *testing.T) {
		wrapper.EXPECT().List(src.MaxCount).Return(taskLists, nil)
		wrapper.EXPECT().Update(id, taskList).Return(nil, xerrors.New("err"))

		op := src.NewTODOOperation(wrapper)

		if _, err := op.UpdateTitleByTODOID(title, title); err == nil {
			t.Error("expected err is not nil")
		}
	})
}
