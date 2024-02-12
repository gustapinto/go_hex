package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gustapinto/go_hex/pkg/httputil"
)

type Logger struct {
	handler http.Handler
}

func WrapWithLogger(handler http.Handler) http.Handler {
	return &Logger{
		handler: handler,
	}
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	l.handler.ServeHTTP(w, r)

	statusCode := r.Context().Value(httputil.StatusCodeKey)

	slog.Info("", "method", r.Method, "code", statusCode, "path", r.URL.Path, "agent", r.UserAgent(), "duration", time.Since(start))
}
