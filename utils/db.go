package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// ListTables lists all the tables in the console
func ListTables() {
	svc := GetDBSession()
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

// HandleDBError returns an appropriate error message based on the error thrown
func HandleDBError(err error) error {
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				return errors.New(string(dynamodb.ErrCodeProvisionedThroughputExceededException) + string(aerr.Error()))
			case dynamodb.ErrCodeResourceNotFoundException:
				return errors.New(string(dynamodb.ErrCodeResourceNotFoundException) + string(aerr.Error()))
			case dynamodb.ErrCodeRequestLimitExceeded:
				return errors.New(string(dynamodb.ErrCodeRequestLimitExceeded) + string(aerr.Error()))
			case dynamodb.ErrCodeInternalServerError:
				return errors.New(string(dynamodb.ErrCodeInternalServerError) + string(aerr.Error()))
			default:
				return errors.New("hmmmm" + string(aerr.Error()))
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			errors.New(string(err.Error()))
		}
	}
	return nil
}

// MakeQuery takes a struct and returns a query
func MakeQuery(in interface{}) (map[string]*dynamodb.AttributeValue, error) {
	return dynamodbattribute.MarshalMap(in)
}

// Get performs a select on the DB
func Get(tableName string, svc *dynamodb.DynamoDB, query map[string]*dynamodb.AttributeValue) (*dynamodb.GetItemOutput, error) {
	return svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       query,
	})
}

// Update performs an update on the DB
func Update(tableName string, svc *dynamodb.DynamoDB, query map[string]*dynamodb.AttributeValue) (*dynamodb.PutItemOutput, error) {
	return svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      query,
	})
}
