package common

import (
    "encoding/base64"
    "github.com/aws/aws-lambda-go/events"
    "time"
)

type APIResponse struct {
    Code int `json:"code"`
    Title string `json:"title"`
    Message string `json:"message"`
    Timestamp time.Time `json:"timestamp"`
}

func PackageResponse(code int, title string, message string) (events.APIGatewayProxyResponse, error) {
    return events.APIGatewayProxyResponse {
        StatusCode: code,
        Body: base64.StdEncoding.EncodeToString([]byte(PrettyJsonString(Response(title, message, code)))),
        IsBase64Encoded: true,
    }, nil
}

func Response(title string, message string, code int) *APIResponse {
    return &APIResponse {
        Code: code,
        Title: title,
        Message: message,
        Timestamp: time.Now(),
    }
}
