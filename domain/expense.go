package domain

type Expense struct {
	ID          int64   `json:"id"`
	User_ID     int64   `json:"user_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	CreatedAt   string  `json:"created_at"`
}
