// having an issue with .aws/config at work, fixed issue at work, now experiencing 400

package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// WcfQuery comment
type WcfQuery struct {
	ID      string `json:"id"`
	Chapter int    `json:"chapter"`
}

func getDBSession() *dynamodb.DynamoDB {
	// lets read something from dynamo db
	sess := session.Must(session.NewSession())
	return dynamodb.New(sess)
}

func listTables(svc *dynamodb.DynamoDB) {
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

func get(svc *dynamodb.DynamoDB, query map[string]*dynamodb.AttributeValue) (*dynamodb.GetItemOutput, error) {
	return svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("wcf"),
		Key:       query,
	})
}

func update(svc *dynamodb.DynamoDB, query map[string]*dynamodb.AttributeValue) (*dynamodb.PutItemOutput, error) {
	return svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("wcf"),
		Item:      query,
	})
}

func main() {
	svc := getDBSession()
	listTables(svc)

	// Read from DB
	chapter1 := WcfQuery{
		ID:      "1",
		Chapter: int(1),
	}

	chapter1Query, err := dynamodbattribute.MarshalMap(chapter1)
	result, err := get(svc, chapter1Query)

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

	fmt.Println("here it is bra", result.Item)
}
