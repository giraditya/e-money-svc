package presentation

type BalanceFetchByUserIDRequest struct {
	UserID string `uri:"userid" validate:"required,numeric"`
}
