package writer

func APIResponse(message string, status bool, data interface{}) Response {
	jsonResponse := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return jsonResponse
}

func APIValidationResponse(message string, status bool, data interface{}, errors interface{}) ValidationResponse {
	jsonResponse := ValidationResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Errors:  errors,
	}

	return jsonResponse
}
