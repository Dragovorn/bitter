package database

import (
	"github.com/aws/aws-sdk-go/aws"
	session2 "github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"main/src/common"
)

var session *session2.Session
var database *dynamo.DB
var constantsProvider common.ConstantsProvider

var initialized = false

func Init(provider common.ConstantsProvider) {
	if initialized {
		return
	}

	constantsProvider = provider
	session = session2.Must(session2.NewSession())
	database = dynamo.New(session, &aws.Config{Region: aws.String("us-east-1")})

	initialized = true
}

func UsersTable() dynamo.Table {
	return Table(constantsProvider.UsersTableName())
}

func Table(name string) dynamo.Table {
	return database.Table(name)
}
