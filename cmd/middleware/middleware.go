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
		return
	}
	switch err.(type) {
	case ErrDbQuery:
		http.Error(w, err.Error(), http.StatusConflict)
	case ErrDbNotSupported:
		http.Error(w, err.Error(), http.StatusConflict)
	case ErrEntityNotFound:
		http.Error(w, err.Error(), http.StatusNotFound)
	default:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	return
}

// ErrDbQuery will be mapped to 409 conflict status
type ErrDbQuery struct {
	Err error
}

func (e ErrDbQuery) Error() string {
	return fmt.Sprintf("%s", e.Err)
}

// ErrDbNotSupported will be mapped to 409 conflict status
type ErrDbNotSupported struct {
	Err error
}

func (e ErrDbNotSupported) Error() string {
	return fmt.Sprintf("%s", e.Err)
}

// ErrEntityNotFound will be mapped to 404 not found status
type ErrEntityNotFound struct {
	Err error
}

func (e ErrEntityNotFound) Error() string {
	return fmt.Sprintf("%s", e.Err)
}
