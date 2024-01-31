package httputil

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