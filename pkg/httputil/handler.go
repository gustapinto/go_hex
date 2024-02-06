package httputil

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type Mux struct {
	mux *http.ServeMux
}

func NewMux() *Mux {
	return &Mux{
		mux: http.NewServeMux(),
	}
}

func (m Mux) Listen(addr string) error {
	return http.ListenAndServe(addr, m.mux)
}

func (m Mux) HandleFunc(method, pattern string, handler http.HandlerFunc) {
	m.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		handler(w, r)
	})
}

func BindJson(w http.ResponseWriter, r *http.Request, target any) error {
	contentType := r.Header.Get(ContentTypeHeader)

	if !strings.Contains(contentType, ContentTypeApplicationJson) {
		err := ErrorResponse{
			Err: ErrorContentTypeNotApplicationJson,
		}

		WriteJson(w, r, http.StatusBadRequest, err)
		return err
	}

	err := json.NewDecoder(r.Body).Decode(&target)
	if err != nil {
		err = ErrorResponse{
			Err: ErrorJsonDecodeFailed,
		}

		WriteJson(w, r, http.StatusBadRequest, err)
		return err
	}

	return nil
}

func WriteJson(w http.ResponseWriter, r *http.Request, successStatusCode int, response any) {
	w.Header().Add(ContentTypeHeader, ContentTypeApplicationJson)

	jsonResponse, err := json.Marshal(&response)
	if err != nil {
		ctx := context.WithValue(r.Context(), StatusCodeKey, http.StatusInternalServerError)
		*r = *(r.WithContext(ctx))

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx := context.WithValue(r.Context(), StatusCodeKey, successStatusCode)
	*r = *(r.WithContext(ctx))

	w.WriteHeader(successStatusCode)
	w.Write(jsonResponse)
}
