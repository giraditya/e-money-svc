package presentation

type TransactionFetchHistoryByUserIDRequest struct {
	UserID uint `uri:"userid" validate:"required,numeric"`
}
