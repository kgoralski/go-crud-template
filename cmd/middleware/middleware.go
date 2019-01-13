package middleware

import (
	"github.com/kgoralski/go-crud-template/internal/banks/domain"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const (
	contentType     = "Content-Type"
	applicationJSON = "application/json"
)

const (
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
	const logFormat = "fatal: %+v\n"
	if strings.Contains(err.Error(), "connection refused") {
		log.Warnf(logFormat, err)
		http.Error(w, DbConnectionFail, http.StatusServiceUnavailable)
		return
	}
	if err.Error() == http.StatusText(400) {
		log.Warnf(logFormat, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	switch err.(type) {
	case domain.ErrDbQuery:
		log.Warnf(logFormat, err.(domain.ErrDbQuery).Err)
		http.Error(w, err.Error(), http.StatusConflict)
	case domain.ErrDbNotSupported:
		log.Warnf(logFormat, err.(domain.ErrDbNotSupported).Err)
		http.Error(w, err.Error(), http.StatusConflict)
	case domain.ErrEntityNotFound:
		log.Warnf(logFormat, err.(domain.ErrEntityNotFound).Err)
		http.Error(w, err.Error(), http.StatusNotFound)
	default:
		log.Warnf(logFormat, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	return
}
