package presentation

type UserFetchByUserIDRequest struct {
	ID uint `json:"id" binding:"required"`
}
