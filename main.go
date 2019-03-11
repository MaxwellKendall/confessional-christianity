// having an issue with .aws/config at work, fixed issue at work, now experiencing 400

package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Chapter comment
type Chapter struct {
	ID      string `json:"id"`
	Whateva string `json:"whateva"`
}

func main() {
	// lets read something from dynamo db
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)
	test := Chapter{
		ID:      "1",
		Whateva: "we",
	}
	key, err := dynamodbattribute.MarshalMap(test)
	if err != nil {
		fmt.Println("YIZO")
	}
	input := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String("test2"),
	}

	result, err := svc.GetItem(input)
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

	fmt.Println(result)

}
