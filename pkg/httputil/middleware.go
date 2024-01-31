package httputil

func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)

		statusCode := r.Context().Value(StatusCodeKey)

		slog.Info("", "method", r.Method, "status_code", statusCode, "path", r.URL.Path, "agent", r.UserAgent())
	}
}
