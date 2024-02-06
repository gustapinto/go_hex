package request

type CreateAccount struct {
	Name         string  `json:"name"`
	InitialValue float64 `json:"initial_value"`
}

type UpdateAccount struct {
	Name         string  `json:"name"`
	CurrentValue float64 `json:"current_value"`
}
