package otuserr

import (
	"encoding/json"
	"net/http"
)

type Const string

func (e Const) Error() string {
	return string(e)
}

const (
	ErrNotFound = Const("not found")
)

func SendSuccess(w http.ResponseWriter, data interface{}) {
	rawJson, err := json.Marshal(data)
	if err != nil {
		SendInternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(rawJson)
}

func SendInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func SendMethodNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("method not allowed"))
}

func SendBadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}

func SendNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}
