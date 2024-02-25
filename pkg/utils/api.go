package utils

type ApiResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func BuildJsonResponse(httpCode int, message string) ApiResponse {
	status := httpCode >= 200 && httpCode <= 210

	return ApiResponse{Status: status, Message: message}
}

func BuildJsonResponseWithData(httpCode int, message string, data any) ApiResponse {
	status := httpCode >= 200 && httpCode <= 210

	return ApiResponse{Status: status, Message: message, Data: data}
}

func BuildJsonResponseWithError(httpCode int, message string, error any) ApiResponse {
	status := httpCode >= 200 && httpCode <= 210

	return ApiResponse{Status: status, Message: message, Error: error}
}

func BuildJsonResponseForValidationError(errors []string) ApiResponse {
	return ApiResponse{Status: false, Message: "Validation error.", Errors: errors}
}
