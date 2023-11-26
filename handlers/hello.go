package handlers

import (
	"io"
	"log"
	"net/http"
)

type HelloHandler struct {
	logger *log.Logger
}

func NewHelloHandler(logger *log.Logger) *HelloHandler {
	return &HelloHandler{logger}
}

func (handler *HelloHandler) GetRoot(w http.ResponseWriter, r *http.Request) {
	handler.logger.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func (handler *HelloHandler) GetHello(w http.ResponseWriter, r *http.Request) {
	handler.logger.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func (handler *HelloHandler) ServeRequestAsResponse(w http.ResponseWriter, r *http.Request) {
	d, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Oops", http.StatusBadRequest)
		return
	}

	w.Write(d)
}
