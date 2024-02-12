package request

type CreateTransaction struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}
