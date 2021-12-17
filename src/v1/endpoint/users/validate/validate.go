package main

import (
	"github.com/gofrs/uuid"
	"main/src/common/bootstrap"
	"main/src/common/request"
	"main/src/common/users"
	v1 "main/src/v1"
	"main/src/v1/entity"
	"strconv"
)

func Run(context *request.LambdaContext) *request.LambdaHTTPResponse {
	var uid uuid.UUID

	uidStr := context.Path["uid"]

	if tUid, err := uuid.FromString(uidStr); err != nil {
		return request.UserError("Malformed User ID", "User ID is malformed!")
	} else {
		uid = tUid
	}

	var user entity.User

	if err := users.Table().Get("uid", uid).One(&user); err != nil {
		return request.DatabaseError(err)
	}

	var code int

	codeStr := context.Query["code"]

	if tCode, err := strconv.Atoi(codeStr); err != nil {
		return request.UserError("Unrecognized Integer", codeStr+" is not a number!")
	} else {
		code = tCode
	}

	return request.Response(200, "Hello", strconv.Itoa(code))
}

func main() {
	bootstrap.All(v1.GetConstants(), Run)
}
