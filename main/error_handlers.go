package main

import (
	"fmt"
	"log"
	"net/http"
)

type httpError struct {
	Err     error
	Message string
	Code    int
}

func (e *httpError) Error() string {
	return fmt.Sprintf("HttpError[%s] Message[%s] Code[%d]", e.Err, e.Message, e.Code)
}

func handleHttpError(w http.ResponseWriter, err error)  {
	switch e := err.(type) {
	case *httpError:
		log.Print(e)
		http.Error(w, e.Message, e.Code)
	default:
		log.Print(err)
		http.Error(w, "UNKNOWN_ERROR", http.StatusInternalServerError)
	}
	return
}