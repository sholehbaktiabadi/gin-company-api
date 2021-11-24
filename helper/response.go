package helper

import (
	"strings"
)

//Response interface
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

//Empty object is used when data doesnt want to be null
type EmptyObj struct{}

// Build Response to return data value success
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

// BuildError Response to return data value failed
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splitterError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Error:   splitterError,
		Data:    data,
	}
	return res
}
