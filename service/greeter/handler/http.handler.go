package handler

import (
	"net/http"
)

type httpHandler struct{}

// NewHTTPHandler returns an instance of `UserServiceHandler`.
func NewHTTPHandler() *httpHandler {
	return &httpHandler{}
}

func (h httpHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("<h1>hello</h1>"))
}
