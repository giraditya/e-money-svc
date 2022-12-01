package presentation

type ErrorValidateResponse struct {
	StatusCode int      `json:"statusCode"`
	Error      []string `json:"error"`
}
