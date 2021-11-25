package v1

import (
    "main/src/common"
    "os"
)

var usersTableName = os.Getenv("USERS_TABLE")
var usernameIndex = os.Getenv("USERNAME_INDEX")
var email = os.Getenv("EMAIL")
var awsRegion = os.Getenv("AWS_REGION")
var instance = Constants{}

func GetConstants() Constants {
    return instance
}

type Constants struct {
    common.ConstantsProvider
}

func (c Constants) UsersTable() string {
    return usersTableName
}

func (c Constants) UsernameIndex() string {
    return usernameIndex
}

func (c Constants) Email() string {
    return email
}

func (c Constants) AWSRegion() string {
    return awsRegion
}