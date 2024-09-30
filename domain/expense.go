package domain

type Expense struct {
	ID          int64   `json:id`
	Name        string  `json:name`
	Description string  `json:description`
	Amount      float64 `json:amount`
}
