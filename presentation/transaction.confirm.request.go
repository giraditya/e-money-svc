package presentation

type TransactionConfirmRequest struct {
	UserID   uint `json:"user_id" binding:"required"`
	BillerID uint `json:"biller_id" binding:"required"`
}
