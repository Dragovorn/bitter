package main

import (
    "encoding/base64"
    "encoding/json"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/gofrs/uuid"
    "golang.org/x/crypto/bcrypt"
    "main/src/common"
    "main/src/common/database"
    "main/src/v1/entity"
)

// Path: /v1/users/register
// Function: Register User

type SubmittedUser struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
    Email string `json:"email" binding:"required"`
}

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
    var submitted SubmittedUser

    bodyRaw, _ := base64.StdEncoding.DecodeString(request.Body)

    body := string(bodyRaw)

    if err := json.Unmarshal([]byte(body), &submitted); err != nil {
        return common.PackageError(400, "Bad Json", "Malformed Input JSON!")
    }

    var potential entity.User

    if err := database.Table().Get("username", submitted.Username).One(&potential); err == nil {
       return common.PackageError(400, "Username Taken", "The username: " + submitted.Username + " is taken!")
    }

    uid, _ := uuid.NewV4()
    passwordHash, _ := bcrypt.GenerateFromPassword([]byte(submitted.Password), 10)

    result := entity.User {
        UID: uid,
        Version: 1,
        Username: submitted.Username,
        PasswordHash: passwordHash,
    }

    if err := database.Table().Put(result).Run(); err != nil {
        return common.DatabaseError()
    }

    return common.PackageResponse(201, "Success", "Successfully created user: " + result.UID.String())
}

func main() {
    database.Init()

    lambda.Start(Handler)
}
