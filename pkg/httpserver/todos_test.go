package httpserver_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	skel "github.com/KentaKudo/goapi-skel"
	"github.com/KentaKudo/goapi-skel/pkg/httpserver"
	"github.com/KentaKudo/goapi-skel/pkg/mock"
	"github.com/gorilla/mux"
)

func TestGetTodos(t *testing.T) {
	inBody, outBody := []skel.Todo{skel.Todo{ID: 0, Title: "test"}}, `{"todos":[{"id":0,"title":"test"}]}`+"\n"
	mock := mock.NewTodoService()
	mock.ListFn = func() ([]skel.Todo, error) {
		return inBody, nil
	}

	sut := (&httpserver.Server{
		Router:      mux.NewRouter(),
		TodoService: mock,
	}).Routes()
	req, _ := http.NewRequest("GET", "/todos", nil)
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("GET /todos: got %v, want %v", got, want)
	}
}

func TestPostTodo(t *testing.T) {
	inBody, outBody := `{"title":"test"}`, &skel.Todo{Title: "test"}
	inID, outID := 123, 123
	mock := mock.NewTodoService()
	mock.CreateFn = func(td *skel.Todo) error {
		if got, want := td, outBody; !reflect.DeepEqual(got, want) {
			t.Errorf("POST /todos: got %v, want %v", got, want)
		}
		td.ID = inID
		return nil
	}

	sut := (&httpserver.Server{
		Router:      mux.NewRouter(),
		TodoService: mock,
	}).Routes()
	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer([]byte(inBody)))
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	if !mock.CreateFnInvoked {
		t.Error("TodoService.Create is not invoked")
	}
	body, _ := ioutil.ReadAll(w.Body)
	if got, want := string(body), fmt.Sprintf(`{"id":%d,"title":"test"}`+"\n", outID); got != want {
		t.Errorf("POST /todos: got %v, want %v", got, want)
	}
}

func TestGetTodo(t *testing.T) {
	inID, outID := 123, 123
	inBody, outBody := &skel.Todo{ID: inID, Title: "test"}, fmt.Sprintf(`{"id":%d,"title":"test"}`, inID)+"\n"
	mock := mock.NewTodoService()
	mock.GetFn = func(id int) (*skel.Todo, error) {
		if got, want := id, outID; got != want {
			t.Errorf("GET /todos/%d: got %v, want %v", inID, got, want)
		}
		return inBody, nil
	}

	sut := (&httpserver.Server{
		Router:      mux.NewRouter(),
		TodoService: mock,
	}).Routes()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/todos/%d", inID), nil)
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	if !mock.GetFnInvoked {
		t.Error("TodoService.Get is not invoked")
	}
	body, _ := ioutil.ReadAll(w.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("GET /todos/%d: got %v, want %v", inID, got, want)
	}
}

func TestDeleteTodo(t *testing.T) {
	inID, outID := 123, 123
	mock := mock.NewTodoService()
	mock.DeleteFn = func(id int) error {
		if got, want := id, outID; got != want {
			t.Errorf("DELETE /todos/%d: got %v, want %v", inID, got, want)
		}
		return nil
	}

	sut := (&httpserver.Server{
		Router:      mux.NewRouter(),
		TodoService: mock,
	}).Routes()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/todos/%d", inID), nil)
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	if !mock.DeleteFnInvoked {
		t.Error("TodoService.Delete is not invoked")
	}
}
