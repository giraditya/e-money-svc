package presentation

type TransactionFetchHistoryByUserIDResponse struct {
	UserID      uint   `json:"userid"`
	Category    string `json:"category"`
	Product     string `json:"product"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Fee         int    `json:"fee"`
	Status      string `json:"status"`
}
