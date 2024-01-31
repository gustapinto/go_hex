package httputil

type ErrorResponse struct {
	Err string `json:"error"`
}

func (e ErrorResponse) Error() string {
	return e.Err
}
