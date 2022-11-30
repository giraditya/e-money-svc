package presentation

type BalanceFetchByUserIDRequest struct {
	UserID string `uri:"userid" binding:"required"`
}
