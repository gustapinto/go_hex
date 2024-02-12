package httputil

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

const (
	ErrorJsonMarshalFailed             = "error.json.marshal.failed"
	ErrorJsonDecodeFailed              = "error.json.decode.failed"
	ErrorContentTypeNotApplicationJson = "error.content.type.not.application.json"

	ContentTypeHeader          = "Content-Type"
	ContentTypeApplicationJson = "application/json"
	StatusCodeKey              = "StatusCode"
)

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

func WriteJson(w http.ResponseWriter, r *http.Request, code int, response any) {
	w.Header().Add(ContentTypeHeader, ContentTypeApplicationJson)

	jsonResponse, err := json.Marshal(&response)
	if err != nil {
		WriteStatusCode(w, r, http.StatusInternalServerError)
		return
	}

	WriteStatusCode(w, r, code)
	w.Write(jsonResponse)
}

func WriteStatusCode(w http.ResponseWriter, r *http.Request, code int) {
	ctx := context.WithValue(r.Context(), StatusCodeKey, code)
	*r = *(r.WithContext(ctx))

	w.WriteHeader(code)
}

func PathValueInt64(w http.ResponseWriter, r *http.Request, name string) (value int64, err error) {
	value, err = strconv.ParseInt(r.PathValue(name), 10, 0)
	if err != nil {
		WriteJson(w, r, http.StatusBadRequest, NewErrorResponse(err))
	}
	return
}
