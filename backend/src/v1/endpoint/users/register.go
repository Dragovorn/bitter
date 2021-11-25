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

    usersTable := database.UsersTable()

    if n, err := usersTable.Scan().Index("usernameIndex").Filter("username = ?", submitted.Username).Count(); n > 0 {
       return common.PackageResponse(400, "Username Taken", "The username: " + submitted.Username + " is taken!")
    } else if err != nil {
        return common.DatabaseError(err)
    }

    uid, _ := uuid.NewV4()
    passwordHash, _ := bcrypt.GenerateFromPassword([]byte(submitted.Password), 10)

    result := entity.User {
        UID: uid,
        Version: 1,
        Username: submitted.Username,
        PasswordHash: passwordHash,
        Email: submitted.Email,
    }

    if err := usersTable.Put(result).Run(); err != nil {
        return common.DatabaseError(err)
    }

    // TODO: Send a validation email to the user

    return common.PackageResponse(201, "Success", "Successfully created user: " + result.UID.String())
}

func main() {
    database.Init(v1.GetConstants())

    lambda.Start(Handler)
}