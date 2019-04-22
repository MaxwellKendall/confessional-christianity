package ccdb

import (
	"errors"
	"strconv"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/MaxwellKendall/confessional-christianity/api"
	"github.com/MaxwellKendall/confessional-christianity/utils"
)

// WcfQuery used to update WCF
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

func unmarshalChapter(dbpayload map[string]*dynamodb.AttributeValue, output *api.WCFChapter) error {
	err := dynamodbattribute.UnmarshalMap(dbpayload, output)
	if err != nil {
		return err
	}
	return nil
}

// GetWcfChapter returns a wcf.Chapter from the db
func GetWcfChapter(chapter int) (api.WCFChapter, error) {
	query, err := utils.MakeQuery(WcfGetQuery{
		ID:      "WCF_" + strconv.Itoa(chapter),
		Chapter: chapter,
	})
	if err != nil {
		return api.WCFChapter{}, errors.New("some error happened when making a query")
	}
	sess := utils.GetDBSession()
	result, err := utils.Get("wcf", sess, query)
	if err != nil {
		return api.WCFChapter{}, utils.HandleDBError(err)
	}
	// TODO: Figure out how to convert this to a struct
	rtrn := api.WCFChapter{}
	unmarshalChapter(result.Item, &rtrn)
	return rtrn, nil
}
