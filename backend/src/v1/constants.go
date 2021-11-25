package v1

import (
    "main/src/common"
    "os"
)

var usersTableName = os.Getenv("USERS_TABLE")
var instance = Constants{}

func GetConstants() Constants {
    return instance
}

type Constants struct {
    common.ConstantsProvider
}

func (c Constants) UsersTableName() string {
    return usersTableName
}
