package presentation

type TransactionFetchHistoryByUserIDResponse struct {
	UserID      uint   `json:"user_id"`
	Category    string `json:"category"`
	Product     string `json:"product"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Fee         int    `json:"fee"`
	Status      string `json:"status"`
}
