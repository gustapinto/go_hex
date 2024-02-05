package httputil

type ErrorResponse struct {
	Err string `json:"error"`
}

func (e ErrorResponse) Error() string {
	return e.Err
}

type CreatedResponse struct {
	ID int64 `json:"id"`
}
