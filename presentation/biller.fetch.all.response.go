package presentation

import "emoney-service/models"

type BillerFetchAllResponse struct {
	Code       int             `json:"code"`
	Status     string          `json:"status"`
	Message    string          `json:"message"`
	BillerData []models.Biller `json:"data"`
}
