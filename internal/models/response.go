package models

import (
	"encoding/json"
	"net/http"
)

// вывод хорошего ответа
type GoodBasicResponse struct {
	Description string `json:"description"        example:"description"`
	Code        int    `json:"code"           example:"status"`
}

// вывод хорошего ответа с машиной
type GoodAdvancedResponse struct {
	Description string `json:"description"        example:"description"`
	Code        int    `json:"code"           example:"status"`
	Cars        []Car  `json:"cars"           example:"...."`
}

// вывод плохого ответа
type ErrorResponse struct {
	Description string `json:"description"`
	Code        int    `json:"code"`
}

func BadResponseSend(w http.ResponseWriter, message string, code int) {
	badResponse := ErrorResponse{
		Description: message,
		Code:        code,
	}
	json.NewEncoder(w).Encode(badResponse)
}

func GoodResponseSend(w http.ResponseWriter, message string, affectedRow int) {
	errResp := GoodBasicResponse{
		Description: message,
		Code:        200,
	}
	json.NewEncoder(w).Encode(errResp)
}
