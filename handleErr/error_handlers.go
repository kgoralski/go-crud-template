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

func HandleErrors(w http.ResponseWriter, err error) {
	log.Print(fmt.Errorf("FATAL: %+v\n", err))
	switch err.Error() {
	case DbQueryFail:
		http.Error(w, err.Error(), http.StatusConflict)
	case DbConnectionFail:
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
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
