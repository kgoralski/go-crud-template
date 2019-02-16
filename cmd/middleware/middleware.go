package middleware

import (
	"net/http"
)

const (
	contentType     = "Content-Type"
	applicationJSON = "application/json"
)

// CommonHeaders to share between packages
func CommonHeaders(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentType, applicationJSON)
		h(w, r)
	}
}
