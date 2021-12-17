package request

import (
	"encoding/base64"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"main/src/common"
	"time"
)

type LambdaHTTPResponse struct {
	Code      int       `json:"code"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Error     error     `json:"-"`
}

func (r *LambdaHTTPResponse) Serialize() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode:      r.Code,
		Body:            base64.StdEncoding.EncodeToString([]byte(common.PrettyJsonString(r))),
		IsBase64Encoded: true,
	}
}

func Response(code int, title string, message string) *LambdaHTTPResponse {
	return &LambdaHTTPResponse{
		Code:      code,
		Title:     title,
		Message:   message,
		Timestamp: time.Now(),
	}
}

func ServerError(message string, err error) *LambdaHTTPResponse {
	response := Response(500, "Internal Server Error!", message)
	response.Error = err

	return response
}

func DatabaseError(err error) *LambdaHTTPResponse {
	return ServerError("Database Error! (uh oh)", err)
}

func UserError(title string, message string) *LambdaHTTPResponse {
	return Response(400, title, message)
}

func UnmarshalJson(str string, do interface{}) *LambdaHTTPResponse {
	if err := json.Unmarshal([]byte(str), &do); err != nil {
		return UserError("Bad Json", "Malformed Input JSON!")
	}

	return nil
}
