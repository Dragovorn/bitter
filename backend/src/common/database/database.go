package database

import (
	"github.com/aws/aws-sdk-go/aws"
	session2 "github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"os"
)

var session *session2.Session
var database *dynamo.DB

var initialized = false

func Init() {
	if initialized {
		return
	}

	session = session2.Must(session2.NewSession())
	database = dynamo.New(session, &aws.Config{Region: aws.String("us-east-1")})

	initialized = true
}

func Table() dynamo.Table {
	return database.Table(os.Getenv("TABLE_NAME"))
}
