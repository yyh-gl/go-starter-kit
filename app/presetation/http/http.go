package http

import (
	"net/http"

	"github.com/yyh-gl/go-starter-kit/app"
)

// WrapHandler adds common header
func WrapHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Add headers for CORS
		switch {
		case app.IsPrd():
			w.Header().Add("Access-Control-Allow-Origin", "https://")
		case app.IsStg():
			w.Header().Add("Access-Control-Allow-Origin", "https://stg.")
		case app.IsDev() || app.IsTest():
			w.Header().Add("Access-Control-Allow-Origin", "http://dev.")
		}
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json;charset=utf-8")

		h.ServeHTTP(w, r)
	}
}

// PreflightHandler is handler of preflight
func PreflightHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
