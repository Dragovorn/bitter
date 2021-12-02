package common

import (
	"github.com/aws/aws-lambda-go/events"
)

func DatabaseError(err error) (events.APIGatewayProxyResponse, error) {
	return Error("Database Error", "An internal database error has occurred! (uh oh)", err)
}

// Error Will properly log an error with a message that technically doesn't matter (for now I hope)
// However it will also email a specific mailing list to notify people of errors, so only use on
// scary errors. If the method returns an error, and you've never had it actually error should probably
// be logged with this method.
func Error(name string, desc string, err error) (events.APIGatewayProxyResponse, error) {
	response, _ := PackageResponse(500, name, desc)

	// TODO: Send an email to a mailing list

	return response, err
}
