package presentation

type UserFetchByUsernameRequest struct {
	Username string `json:"username" binding:"required"`
}
