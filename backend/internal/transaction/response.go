package transaction

type TransactionResponse struct {
	ID              uint    `json:"id"`
	UserID          uint    `json:"user_id"`
	CategoryID      uint    `json:"category_id"`
	CategoryName    string  `json:"category_name"`
	Amount          float64 `json:"amount"`
	Description     string  `json:"description"`
	TransactionDate string  `json:"transaction_date"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}
