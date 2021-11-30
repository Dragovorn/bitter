package main

import (
    "github.com/aws/aws-lambda-go/events"
    "main/src/common"
    "main/src/common/bootstrap"
    v1 "main/src/v1"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
    return common.PackageResponse(200, "Hello", "World!")
}

func main() {
    bootstrap.All(v1.GetConstants(), Handler)
}
