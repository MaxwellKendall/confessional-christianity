package db

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Get acts on a table
func Get(svc *dynamodb.DynamoDB, query map[string]*dynamodb.AttributeValue) (*dynamodb.GetItemOutput, error) {
	return svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("wcf"),
		Key:       query,
	})
}

// Update acts on a table
func Update(svc *dynamodb.DynamoDB, query map[string]*dynamodb.AttributeValue) (*dynamodb.PutItemOutput, error) {
	return svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("wcf"),
		Item:      query,
	})
}

// HandleDBError handles an error from the db
func HandleDBError(err error) {
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println("hmmmm", aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
}

// ListTables lists the tables in the DB
func ListTables(svc *dynamodb.DynamoDB) {
	result, err := svc.ListTables(&dynamodb.ListTablesInput{})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Tables:")
	fmt.Println("")

	for _, n := range result.TableNames {
		fmt.Println(*n)
	}
}

// MakeQuery makes a query from a map
func MakeQuery(in interface{}) (map[string]*dynamodb.AttributeValue, error) {
	return dynamodbattribute.MarshalMap(in)
}

// GetDBSession gets a session from the db
func GetDBSession() *dynamodb.DynamoDB {
	// lets read something from dynamo db
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile: "ccdev",
	}))
	return dynamodb.New(sess)
}
