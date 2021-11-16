package common

import (
    "github.com/aws/aws-lambda-go/events"
    "time"
)

type APIResponse struct {
    Title string `json:"title"`
    Message string `json:"message"`
    Timestamp time.Time `json:"timestamp"`
}

func PackageResponse(code int, title string, message string) (events.APIGatewayProxyResponse, error) {
    return events.APIGatewayProxyResponse {
        StatusCode: code,
        Body: PrettyJsonString(Response(title, message)),
    }, nil
}

func Response(title string, message string) *APIResponse {
    return &APIResponse {
        Title: title,
        Message: message,
        Timestamp: time.Now(),
    }
}
