package response

import "github.com/api-abc/internal-api-module/model/domain"

type BodyResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    domain.Data `json:"data,omitempty"`
}

type BodyResponseGet struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []domain.Data `json:"data,omitempty"`
}
