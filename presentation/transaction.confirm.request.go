package presentation

type TransactionConfirmRequest struct {
	UserID   uint `json:"userid" validate:"required,numeric"`
	BillerID uint `json:"billerid" validate:"required,numeric"`
}
