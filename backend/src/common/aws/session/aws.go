package session

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "main/src/common"
)

var awsSession *session.Session

func Init() {
    awsSession = session.Must(session.NewSession(&aws.Config{Region: aws.String(common.Constants().AWSRegion())}))
}

func Session() *session.Session {
    return awsSession
}
