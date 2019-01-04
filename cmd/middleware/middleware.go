package middleware

import (
	"fmt"
	log "github.com/sirupsen/logrus"
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
func CommonHeaders(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentType, applicationJSON)
		h(w, r)
	}
}

// HandleErrors , DB errors to Rest mapper
func HandleErrors(w http.ResponseWriter, err error) {
	log.Warn(fmt.Errorf("fatal: %+v", err))
	if strings.Contains(err.Error(), "connection refused") {
		http.Error(w, DbConnectionFail, http.StatusServiceUnavailable)
		return
	}
	if err.Error() == http.StatusText(400) {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	switch err.(type) {
	case DbQueryError:
		http.Error(w, err.Error(), http.StatusConflict)
	case DbNotSupportedError:
		http.Error(w, err.Error(), http.StatusConflict)
	case EntityNotFoundError:
		http.Error(w, err.Error(), http.StatusNotFound)
	default:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	return
}

// DbQueryError will be mapped to 409 conflict status
type DbQueryError struct {
	Err error
}

func (e DbQueryError) Error() string {
	return fmt.Sprintf("%s", e.Err)
}

// DbNotSupportedError will be mapped to 409 conflict status
type DbNotSupportedError struct {
	Err error
}

func (e DbNotSupportedError) Error() string {
	return fmt.Sprintf("%s", e.Err)
}

// EntityNotFoundError will be mapped to 404 not found status
type EntityNotFoundError struct {
	Err error
}

func (e EntityNotFoundError) Error() string {
	return fmt.Sprintf("%s", e.Err)
}
