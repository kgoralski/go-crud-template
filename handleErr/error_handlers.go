package handleErr

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const (
	// DbQueryFail represents db query failures
	DbQueryFail = "DB_QUERY_FAIL"
	// DbNotSupported represents db not supported operation
	DbNotSupported = "DB_NOT_SUPPORTED"
	// EntityNotExist represents error that entity doesn't exist in db
	EntityNotExist = "ENTITY_NOT_EXIST"
	// DbConnectionFail represents that application couldn't connect to db
	DbConnectionFail = "DB_CONNECTION_FAIL"
)

// HandleErrors helps handling errors
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
