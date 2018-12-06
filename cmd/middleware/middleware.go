package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	contentType     = "Content-Type"
	applicationJSON = "application/json"
)

const (
	// DbQueryFail represents DB query failures
	DbQueryFail = "DB_QUERY_FAIL"
	// DbNotSupported represents DB not supported operation
	DbNotSupported = "DB_NOT_SUPPORTED"
	// EntityNotExist represents error that entity doesn't exist in DB
	EntityNotExist = "ENTITY_NOT_EXIST"
	// DbConnectionFail represents that application couldn't connect to DB
	DbConnectionFail = "DB_CONNECTION_FAIL"
)

// CommonHeaders to share between packages
func CommonHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentType, applicationJSON)
		fn(w, r)
	}
}

// HandleErrors , DB errors to Rest mapper
func HandleErrors(w http.ResponseWriter, err error) {
	log.Print(fmt.Errorf("fatal: %+v", err))
	if strings.Contains(err.Error(), "connection refused") {
		http.Error(w, DbConnectionFail, http.StatusServiceUnavailable)
		return
	}
	switch err.Error() {
	case DbQueryFail:
		http.Error(w, err.Error(), http.StatusConflict)
	case DbNotSupported:
		http.Error(w, err.Error(), http.StatusConflict)
	case EntityNotExist:
		http.Error(w, err.Error(), http.StatusNotFound)
	case http.StatusText(400):
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	return
}
