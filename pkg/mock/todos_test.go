package mock_test

import (
	"reflect"
	"testing"

	skel "github.com/KentaKudo/goapi-skel"
	"github.com/KentaKudo/goapi-skel/pkg/mock"
)

func TestGet(t *testing.T) {
	mockID := 123
	mockTodo := &skel.Todo{ID: 456, Title: "test"}
	sut := mock.NewTodoService()
	sut.GetFn = func(id int) (*skel.Todo, error) {
		if got, want := id, mockID; got != want {
			t.Errorf("Get(%d): got %v, want %v", mockID, got, want)
		}
		return mockTodo, nil
	}

	out, err := sut.Get(mockID)
	if !sut.GetFnInvoked {
		t.Error("GetFnInvoked returns false")
	}
	if err != nil {
		t.Errorf("Get(%d): unexpected error %v", mockID, err)
	}
	if got, want := out, mockTodo; !reflect.DeepEqual(got, want) {
		t.Errorf("Get(%d): got %v, want %v", mockID, got, want)
	}
}

func TestList(t *testing.T) {
	mockTodos := []skel.Todo{
		skel.Todo{ID: 123, Title: "test1"},
		skel.Todo{ID: 456, Title: "test2"},
	}
	sut := mock.NewTodoService()
	sut.ListFn = func() ([]skel.Todo, error) {
		return mockTodos, nil
	}

	out, err := sut.List()
	if !sut.ListFnInvoked {
		t.Error("ListFnInvoked returns false")
	}
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if got, want := out, mockTodos; !reflect.DeepEqual(got, want) {
		t.Errorf("List(): got %v, want %v", got, want)
	}
}

func TestCreate(t *testing.T) {
	mockTodo := &skel.Todo{Title: "test"}
	sut := mock.NewTodoService()
	sut.CreateFn = func(td *skel.Todo) error {
		if got, want := td, mockTodo; !reflect.DeepEqual(got, want) {
			t.Errorf("Create(%v): got %v, want %v", mockTodo, got, want)
		}
		return nil
	}

	if err := sut.Create(mockTodo); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !sut.CreateFnInvoked {
		t.Error("CreateFnIvoked returns false")
	}
}

func TestDelete(t *testing.T) {
	mockID := 123
	sut := mock.NewTodoService()
	sut.DeleteFn = func(id int) error {
		if got, want := id, mockID; got != want {
			t.Errorf("Delete(%v): got %v, want %v", mockID, got, want)
		}
		return nil
	}

	if err := sut.Delete(mockID); err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if !sut.DeleteFnInvoked {
		t.Error("DeleteFnInvoked returns false")
	}
}
