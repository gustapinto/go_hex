package httputil

import (
	"net/http"
	"log/slog"
)

func Log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)

		statusCode := r.Context().Value(StatusCodeKey)

		slog.Info("", "method", r.Method, "status_code", statusCode, "path", r.URL.Path, "agent", r.UserAgent())
	}
}
