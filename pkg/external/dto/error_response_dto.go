package dto

type ErrorResponseDTO struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
