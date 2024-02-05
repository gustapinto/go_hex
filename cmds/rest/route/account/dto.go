package account

type CreateAccountRequest struct {
	Name         string  `json:"name"`
	InitialValue float64 `json:"initial_value"`
}
