// having an issue with .aws/config at work, fixed issue at work, now experiencing 400

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/MaxwellKendall/confessional-christianity/scripts"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

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

func handleDBError(err error) {
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

func makeQuery(in interface{}) (map[string]*dynamodb.AttributeValue, error) {
	return dynamodbattribute.MarshalMap(in)
}

func main() {
	svc := getDBSession()
	wcf := wcf.ImportWCF()

	for _, chap := range wcf {
		parsedChapter := WcfGetQuery{
			ID:      "WCF_" + strconv.Itoa(chap.Number),
			Chapter: chap.Number,
		}
		query, err := makeQuery(parsedChapter)
		if err != nil {
			fmt.Println("**ERROR: ", err)
		}
		result, err := get(svc, query)
		handleDBError(err)
		fmt.Println(result)
	}

	// fmt.Println("Result", output)
}
