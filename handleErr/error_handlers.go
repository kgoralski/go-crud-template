package handleErr

import (
	"fmt"
	"log"
	"net/http"
)

const (
	DbQueryFail      = "DB_QUERY_FAIL"
	DbNotSupported   = "DB_NOT_SUPPORTED"
	EntityNotExist   = "ENTITY_NOT_EXIST"
	DbConnectionFail = "DB_CONNECTION_FAIL"
)

type HTTPError struct {
	Err     error
	Message string
	Code    int
}

type DbError struct {
	Err     error
	Message string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HttpError[%s] Message[%s] Code[%d]", e.Err, e.Message, e.Code)
}

func (e *DbError) Error() string {
	return fmt.Sprintf("DbError[%s] Message[%s]", e.Err, e.Message)
}

func HandleErrors(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case *HTTPError:
		log.Print(e)
		http.Error(w, e.Message, e.Code)
	case *DbError:
		log.Print(e)
		switch e.Message {
		case DbQueryFail:
			http.Error(w, DbQueryFail, http.StatusConflict)
		case DbConnectionFail:
			http.Error(w, DbConnectionFail, http.StatusServiceUnavailable)
		case DbNotSupported:
			http.Error(w, DbNotSupported, http.StatusConflict)
		case EntityNotExist:
			http.Error(w, EntityNotExist, http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	default:
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	return
}
