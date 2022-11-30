package presentation

type BalanceFetchByUserIDResponse struct {
	UserID  uint `json:"user_id"`
	Balance int  `json:"balance"`
}
