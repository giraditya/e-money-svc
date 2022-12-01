package presentation

type UserFetchByUserIDRequest struct {
	ID uint `json:"id" validate:"required,numeric"`
}
