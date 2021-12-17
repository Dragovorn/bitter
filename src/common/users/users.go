package users

import (
	"github.com/guregu/dynamo"
	"main/src/common"
	"main/src/common/aws/database"
)

func New(user interface{}) error {
	return Table().Put(database.Serialize(user)).Run()
}

func ByUsername() *dynamo.Scan {
	return Table().Scan().Index(common.Constants().UsernameIndex())
}

func Table() dynamo.Table {
	return database.UsersTable()
}
