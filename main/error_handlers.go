package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	dbQueryFail      = "DB_QUERY_FAIL"
	dbNotSupported   = "DB_NOT_SUPPORTED"
	entityNotExist   = "ENTITY_NOT_EXIST"
	dbConnectionFail = "DB_CONNECTION_FAIL"
)

type httpError struct {
	Err     error
	Message string
	Code    int
}

type dbError struct {
	Err     error
	Message string
}

func (e *httpError) Error() string {
	return fmt.Sprintf("HttpError[%s] Message[%s] Code[%d]", e.Err, e.Message, e.Code)
}

func (e *dbError) Error() string {
	return fmt.Sprintf("DbError[%s] Message[%s]", e.Err, e.Message)
}

func handleErrors(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case *httpError:
		log.Print(e)
		http.Error(w, e.Message, e.Code)
	case *dbError:
		log.Print(e)
		switch e.Message {
		case dbQueryFail:
			http.Error(w, dbQueryFail, http.StatusConflict)
		case dbConnectionFail:
			http.Error(w, dbConnectionFail, http.StatusServiceUnavailable)
		case dbNotSupported:
			http.Error(w, dbNotSupported, http.StatusConflict)
		case entityNotExist:
			http.Error(w, entityNotExist, http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	default:
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	return
}
