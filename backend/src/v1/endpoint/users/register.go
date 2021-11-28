package main

import (
    "encoding/base64"
    "encoding/json"
    "github.com/aws/aws-lambda-go/events"
    "main/src/common"
    "main/src/common/aws/email"
    "main/src/common/bootstrap"
    "main/src/common/users"
    v1 "main/src/v1"
    "main/src/v1/entity"
)

// Path: /v1/users/register
// Function Description: Registers a new User

type SubmittedUser struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Email string `json:"email"`
}

func (su *SubmittedUser) IsValid() bool {
    return su != nil && su.Username != "" && su.Password != "" && su.Email != ""
}

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
    var submitted SubmittedUser

    bodyRaw, _ := base64.StdEncoding.DecodeString(request.Body)

    body := string(bodyRaw)

    if err := json.Unmarshal([]byte(body), &submitted); err != nil {
        return common.PackageResponse(400, "Bad Json", "Malformed Input JSON!")
    }

    if !submitted.IsValid() {
        return common.PackageResponse(400, "Bad Input", "Missing required data, username, password, email!")
    }

    // TODO: Validate username & password against some policy I come up with

    if n, err := users.ByUsername().Filter("username = ?", submitted.Username).Count(); n > 0 {
       return common.PackageResponse(400, "Username Taken", "The username: " + submitted.Username + " is taken!")
    } else if err != nil {
        return common.DatabaseError(err)
    }

    result := entity.NewUser(submitted.Username, submitted.Email, submitted.Password)

    // TODO: Make this a validation email
    message := email.TextMessage("Hello, World!!!", "Hello " + result.Username + "!")

    if _, err := email.Send(email.To(result.Email), message); err != nil {
        return common.Error("Email failed!", "Could not send email!", err)
    }

    if err := users.New(result); err != nil {
        return common.DatabaseError(err)
    }

    return common.PackageResponse(201, "Success", "Successfully created user: " + result.UID.String())
}

func main() {
    bootstrap.All(v1.GetConstants(), Handler)
}
