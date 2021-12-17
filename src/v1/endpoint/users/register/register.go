package main

import (
	"main/src/common/aws/email"
	"main/src/common/bootstrap"
	"main/src/common/request"
	"main/src/common/users"
	"main/src/common/users/validation"
	v1 "main/src/v1"
	"main/src/v1/entity"
	"strconv"
)

// Path: /v1/users/register
// Function Description: Registers a new User

type SubmittedUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (su *SubmittedUser) IsValid() bool {
	return su != nil && su.Username != "" && su.Password != "" && su.Email != ""
}

func Run(context *request.LambdaContext) *request.LambdaHTTPResponse {
	var submitted SubmittedUser

	if err := context.UnmarshalBody(&submitted); err != nil {
		return err
	}

	if !submitted.IsValid() {
		return request.UserError("Bad Input", "Missing required data, username, password, email!")
	}

	// TODO: Validate username & password against some policy I come up with

	if n, err := users.ByUsername().Filter("username = ?", submitted.Username).Count(); n > 0 {
		return request.UserError("Username Taken", "The username: "+submitted.Username+" is taken!")
	} else if err != nil {
		return request.DatabaseError(err)
	}

	result := entity.NewUser(submitted.Username, submitted.Email, submitted.Password)
	code := validation.NewCode(result.UID)

	if err := validation.New(code); err != nil {
		return request.DatabaseError(err)
	}

	message := email.TextMessage("Your validation code is: "+strconv.Itoa(code.Code), "Hello "+result.Username+"!")

	if _, err := email.Send(email.To(result.Email), message); err != nil {
		return request.ServerError("Could not send email!", err)
	}

	if err := users.New(result); err != nil {
		return request.DatabaseError(err)
	}

	return request.Response(201, "Success", "Successfully created user: "+result.UID.String())
}

func main() {
	bootstrap.All(v1.GetConstants(), Run)
}
