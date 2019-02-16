package skel_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	skel "github.com/KentaKudo/goapi-skel"
	"github.com/KentaKudo/goapi-skel/mock"
)

func TestGetTodos(t *testing.T) {
	inBody, outBody := []skel.Todo{skel.Todo{Title: "test"}}, `{"todos":[{"title":"test"}]}`+"\n"
	mock := mock.NewTodoService()
	mock.ListFn = func() ([]skel.Todo, error) {
		return inBody, nil
	}

	sut := skel.New(mock).Routes()
	req, _ := http.NewRequest("GET", "/todos", nil)
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("GET /todos: got %v, want %v", got, want)
	}
}
