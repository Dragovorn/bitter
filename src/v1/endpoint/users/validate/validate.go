package main

import (
	"github.com/gofrs/uuid"
	"main/src/common/bootstrap"
	"main/src/common/request"
	"main/src/common/users"
	"main/src/common/users/validation"
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

	codeStr := context.Query["code"]

	if _, err := strconv.Atoi(codeStr); err != nil {
		return request.UserError("Unrecognized Integer", codeStr+" is not a number!")
	}

	code := validation.Code{
		Code: codeStr,
	}

	// TODO: Database error here: "Provided key element does not match the schema"
	if err := validation.Table().Get("code", code.Code).One(&code); err != nil {
		panic(err) // TODO: Handle more gracefully, just experiment with this error
	}

	// debt, this code block can probably be factored into the db query with some work
	// this might also become necessary if codes collide.
	if code.UserId != uid {
		return request.UserError("Unrecognized Code", "Your validation code is unrecognized!")
	}

	// TODO: Some database shit

	return request.Response(200, "Validated!", "You have successfully validated your account!")
}

func main() {
	bootstrap.All(v1.GetConstants(), Run)
}
