package response

type ErrorResponse struct {
	Err string `json:"error"`
}

func (e ErrorResponse) Error() string {
	return e.Err
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Err: err.Error(),
	}
}

type CreatedResponse struct {
	ID int64 `json:"id"`
}

func NewCreatedResponse(id int64) CreatedResponse {
	return CreatedResponse{
		ID: id,
	}
}
