package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// WcfPutQuery used to update WCF
type WcfQuery struct {
	ID         string            `json:"id"`
	Chapter    int               `json:"chapter"`
	Title      string            `json:"title"`
	Paragraphs []string          `json:"paragraphs"`
	Proofs     map[string]string `json:"proofs"`
}

// WcfGetQuery used to update WCF
type WcfGetQuery struct {
	ID      string `json:"id"`
	Chapter int    `json:"chapter"`
}

// Get performs a select on the DB
func Get(svc *dynamodb.DynamoDB, query map[string]*dynamodb.AttributeValue) (*dynamodb.GetItemOutput, error) {
	return svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("wcf"),
		Key:       query,
	})
}

// Update performs an update on the DB
func Update(svc *dynamodb.DynamoDB, query map[string]*dynamodb.AttributeValue) (*dynamodb.PutItemOutput, error) {
	return svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("wcf"),
		Item:      query,
	})
}
