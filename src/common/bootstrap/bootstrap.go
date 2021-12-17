package bootstrap

import (
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"main/src/common"
	"main/src/common/aws/database"
	"main/src/common/aws/email"
	"main/src/common/aws/session"
	"main/src/common/request"
)

var runnerFunc func(ctx *request.LambdaContext) *request.LambdaHTTPResponse

func Bootstrap(req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var body string
	var bodyRaw []byte

	if req.IsBase64Encoded {
		bodyRaw, _ = base64.StdEncoding.DecodeString(req.Body)

		body = string(bodyRaw)
	} else {
		body = req.Body
		bodyRaw = []byte(body)
	}

	ctx := &request.LambdaContext{
		BodyRaw:   bodyRaw,
		Body:      body,
		Path:      req.PathParameters,
		Query:     req.QueryStringParameters,
		RawPath:   req.RawPath,
		RawQuery:  req.RawQueryString,
		WasBase64: req.IsBase64Encoded,
	}

	response := runnerFunc(ctx)

	if response.Error != nil {
		fmt.Println("------")
		fmt.Println("Encountered Error:")
		fmt.Println(response.Error.Error())
		fmt.Println("------")
	}

	return response.Serialize(), nil
}

func All(provider common.ConstantsProvider, runner func(ctx *request.LambdaContext) *request.LambdaHTTPResponse) {
	runnerFunc = runner

	common.Init(provider)
	session.Init()
	database.Init()
	email.Init()

	lambda.Start(Bootstrap)
}
