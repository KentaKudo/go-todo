package httpserver

import (
	"encoding/json"
	"io"
	"net/http"
)

func decode(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func encode(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func fail(w http.ResponseWriter, ierr *Error) {
	w.WriteHeader(ierr.Code)
	w.Write([]byte(ierr.Error()))
}
