package httpserver_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
