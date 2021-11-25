package database

import (
	"github.com/guregu/dynamo"
	"main/src/common"
	"main/src/common/aws/session"
)

var database *dynamo.DB

type Serializable interface {
    Serialize() interface {}
}

func Init() {
	database = dynamo.New(session.Session())
}

func UsersTable() dynamo.Table {
	return Table(common.Constants().UsersTable())
}

func Table(name string) dynamo.Table {
	return database.Table(name)
}
