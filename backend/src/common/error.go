package common

import (
	"github.com/aws/aws-lambda-go/events"
	"time"
)

type APIError struct {
	Code int `json:"code"`
	Name string `json:"name"`
	Message string `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func DatabaseError() (events.APIGatewayProxyResponse, error) {
	return PackageError(500, "Database Error", "An internal database error has occurred! (uh oh)")
}

func PackageError(code int, name string, message string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse {
		StatusCode: code,
		Body: PrettyJsonString(Error(code, name, message)),
	}, nil
}

func Error(code int, name string, message string) *APIError {
	return &APIError{
		Code: code,
		Name: name,
		Message: message,
		Timestamp: time.Now(),
	}
}
