package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// TODO: is code needed?
func APIResponse(message string, code int, status string, data interface{}) Response {
	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return response
}

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func FormatNotFoundError(status string, message string, data interface{}) Response {
	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return response
}

func FormatBadRequest(data interface{}) Response {
	response := Response{
		Status:  "Bad Request",
		Message: "title cannot be null",
		Data:    data,
	}

	return response
}
