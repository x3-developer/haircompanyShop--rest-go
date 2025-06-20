package response

import (
	"encoding/json"
	"net/http"
)

type ErrorField struct {
	Field     string `json:"field"`
	ErrorCode string `json:"errorCode"`
}

type successResponse struct {
	IsSuccess bool `json:"isSuccess"`
	Data      any  `json:"data,omitempty"`
}

type errorResponse struct {
	IsSuccess bool         `json:"isSuccess"`
	Message   string       `json:"message"`
	ErrorCode string       `json:"errorCode"`
	Fields    []ErrorField `json:"fields,omitempty"`
}

func NewErrorField(field, errorCode string) ErrorField {
	return ErrorField{
		Field:     field,
		ErrorCode: errorCode,
	}
}

func SendSuccess(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := &successResponse{
		IsSuccess: true,
		Data:      data,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode success response", http.StatusInternalServerError)
	}
}

func SendError(w http.ResponseWriter, statusCode int, message string, errorCode ErrorCode) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := &errorResponse{
		IsSuccess: false,
		Message:   message,
		ErrorCode: string(errorCode),
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
	}
}

func SendValidationError(w http.ResponseWriter, statusCode int, message string, errorCode ErrorCode, fields []ErrorField) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := &errorResponse{
		IsSuccess: false,
		Message:   message,
		ErrorCode: string(errorCode),
		Fields:    fields,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode validation error response", http.StatusInternalServerError)
	}
}
