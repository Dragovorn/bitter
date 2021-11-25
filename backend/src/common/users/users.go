package users

import (
    "github.com/guregu/dynamo"
    "main/src/common"
    "main/src/common/aws/database"
)

func New(user interface {}) error {
    var val interface {}

    switch user.(type) {
    default:
        val = user
    case database.Serializable:
        val = user.(database.Serializable).Serialize()
    }

    return Table().Put(val).Run()
}

func ByUsername() *dynamo.Scan {
    return Table().Scan().Index(common.Constants().UsernameIndex())
}

func Table() dynamo.Table {
    return database.UsersTable()
}
