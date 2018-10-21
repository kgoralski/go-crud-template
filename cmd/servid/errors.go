package servid

import (
	"fmt"
	"github.com/kgoralski/go-crud-template/internal/banks"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func handleErrors(w http.ResponseWriter, err error) {
	log.Print(fmt.Errorf("fatal: %+v", err))
	if strings.Contains(err.Error(), "connection refused") {
		http.Error(w, banks.DbConnectionFail, http.StatusServiceUnavailable)
		return
	}
	switch err.Error() {
	case banks.DbQueryFail:
		http.Error(w, err.Error(), http.StatusConflict)
	case banks.DbNotSupported:
		http.Error(w, err.Error(), http.StatusConflict)
	case banks.EntityNotExist:
		http.Error(w, err.Error(), http.StatusNotFound)
	case http.StatusText(400):
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	return

}
