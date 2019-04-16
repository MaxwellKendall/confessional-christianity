package utils

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func getDBSession() *dynamodb.DynamoDB {
	// lets read something from dynamo db
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile: "ccdev",
	}))

	return dynamodb.New(sess)
}

// GetConfig returns config object with all the necessary tools for the service
func GetConfig() interface{} {
	// 1. return value should be a user defined type, right now it's just an interface
	// 2. return db session here
	rtrn := struct{
		// session: getDbSession()
	}
}
