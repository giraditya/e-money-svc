package presentation

type ErrorResponse struct {
	StatusCode int         `json:"statusCode"`
	Error      interface{} `json:"error"`
}
