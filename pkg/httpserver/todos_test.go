package httpserver_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	app "github.com/KentaKudo/goapi-skel/pkg"
	"github.com/KentaKudo/goapi-skel/pkg/mock"
)

func TestGetTodos(t *testing.T) {
	inBody, outBody := []app.Todo{app.Todo{Title: "test"}}, `{"todos":[{"title":"test"}]}`+"\n"
	mock := mock.NewTodoService()
	mock.ListFn = func() ([]app.Todo, error) {
		return inBody, nil
	}

	sut := app.New(mock).Routes()
	req, _ := http.NewRequest("GET", "/todos", nil)
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("GET /todos: got %v, want %v", got, want)
	}
}
