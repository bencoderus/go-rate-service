package utils

type ApiResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func BuildJsonResponse(httpCode int, message string) ApiResponse {
	status := httpCode >= 200 && httpCode <= 210

	return ApiResponse{Status: status, Message: message}
}

func BuildJsonResponseWithData(httpCode int, message string, data interface{}) ApiResponse {
	status := httpCode >= 200 && httpCode <= 210

	return ApiResponse{Status: status, Message: message, Data: data}
}
