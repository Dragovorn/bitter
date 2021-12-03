package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/gofrs/uuid"
	"main/src/common"
	"main/src/common/bootstrap"
	"main/src/common/users"
	v1 "main/src/v1"
	"main/src/v1/entity"
	"strconv"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var uid uuid.UUID

	uidStr := request.PathParameters["uid"]

	if tUid, err := uuid.FromString(uidStr); err != nil {
		return common.PackageResponse(400, "Malformed User ID", "User ID is malformed!")
	} else {
		uid = tUid
	}

	var user entity.User

	if err := users.Table().Get("uid", uid).One(&user); err != nil {
		panic(err)
	}

	var code int

	codeStr := request.QueryStringParameters["code"]

	if tCode, err := strconv.Atoi(codeStr); err != nil {
		return common.PackageResponse(400, "Unrecognized Integer", codeStr+" is not a number!")
	} else {
		code = tCode
	}

	return common.PackageResponse(200, "Hello", strconv.Itoa(code))
}

func main() {
	bootstrap.All(v1.GetConstants(), Handler)
}
