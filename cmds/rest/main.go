package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

const (
	ERROR_JSON_MARSHAL_FAILED                                  = "error.json.marshal.failed"
	ERROR_JSON_DECODE_FAILED                                   = "error.json.decode.failed"
	ERROR_JSON_DECODE_FAILED_CONTENT_TYPE_NOT_APPLICATION_JSON = "error.json.decode.failed.content.type.not.application.json"

	CONTENT_TYPE_HEADER_KEY       = "Content-Type"
	CONTENT_TYPE_APPLICATION_JSON = "application/json"

	SERVER_ADDRESS = "0.0.0.0:8080"

	REQUEST_CONTEXT_STATUS_CODE_KEY = "STATUS_CODE"
)

type PingResponse struct {
	Ping string `json:"ping"`
}

type ErrorResponse struct {
	Err string `json:"error"`
}

func (e ErrorResponse) Error() string {
	return e.Err
}

func BindJson(w http.ResponseWriter, r *http.Request, target any) error {
	contentType := r.Header.Get(CONTENT_TYPE_HEADER_KEY)

	if !strings.Contains(contentType, CONTENT_TYPE_APPLICATION_JSON) {
		err := ErrorResponse{
			Err: ERROR_JSON_DECODE_FAILED_CONTENT_TYPE_NOT_APPLICATION_JSON,
		}

		WriteJson(w, r, http.StatusBadRequest, err)
		return err
	}

	err := json.NewDecoder(r.Body).Decode(&target)
	if err != nil {
		err = ErrorResponse{
			Err: ERROR_JSON_DECODE_FAILED,
		}

		WriteJson(w, r, http.StatusBadRequest, err)
		return err
	}

	return nil
}

func WriteJson(w http.ResponseWriter, r *http.Request, statusCode int, response any) {
	w.Header().Add(CONTENT_TYPE_HEADER_KEY, CONTENT_TYPE_APPLICATION_JSON)

	jsonResponse, err := json.Marshal(&response)
	if err != nil {
		*r = *(r.WithContext(context.WithValue(r.Context(), REQUEST_CONTEXT_STATUS_CODE_KEY, http.StatusInternalServerError)))

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	*r = *(r.WithContext(context.WithValue(r.Context(), REQUEST_CONTEXT_STATUS_CODE_KEY, statusCode)))

	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}

func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)

		statusCode := r.Context().Value(REQUEST_CONTEXT_STATUS_CODE_KEY)

		slog.Info("", "method", r.Method, "status_code", statusCode, "path", r.URL.Path, "agent", r.UserAgent())
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/ping", LogRequest(func(w http.ResponseWriter, r *http.Request) {
		WriteJson(w, r, http.StatusOK, &PingResponse{
			Ping: "pong",
		})
	}))

	slog.Info("Starting HTTP server", "address", SERVER_ADDRESS)

	if err := http.ListenAndServe(SERVER_ADDRESS, mux); err != nil {
		slog.Error("Failed to start http server", "error", err.Error())
		os.Exit(1)
	}
}
