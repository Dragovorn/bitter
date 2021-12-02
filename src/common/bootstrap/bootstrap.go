package bootstrap

import (
    "github.com/aws/aws-lambda-go/lambda"
    "main/src/common"
    "main/src/common/aws/database"
    "main/src/common/aws/email"
    "main/src/common/aws/session"
)

func All(provider common.ConstantsProvider, lambdaHandler interface {}) {
    common.Init(provider)
    session.Init()
    database.Init()
    email.Init()

    lambda.Start(lambdaHandler)
}
