package handleErr

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	DbQueryFail      = "DB_QUERY_FAIL"
	DbNotSupported   = "DB_NOT_SUPPORTED"
	EntityNotExist   = "ENTITY_NOT_EXIST"
	DbConnectionFail = "DB_CONNECTION_FAIL"
)

func HandleErrors(w http.ResponseWriter, err error) {
	log.Print(fmt.Errorf("FATAL: %+v\n", err))
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
