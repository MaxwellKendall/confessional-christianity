package ccdb

import (
	"errors"
	"strconv"

	"github.com/MaxwellKendall/confessional-christianity/impl/api"
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

// GetWcfChapter returns a wcf.Chapter from the db
func GetWcfChapter(chapter int) (interface{}, error) {
	query, err := utils.MakeQuery(WcfGetQuery{
		ID:      "WCF_" + strconv.Itoa(chapter),
		Chapter: chapter,
	})
	if err != nil {
		return api.Chapter{}, errors.New("some error happened when making a query")
	}
	svc := utils.GetDBSession()
	result, err := utils.Get("wcf", svc, query)
	if err != nil {
		return api.Chapter{}, utils.HandleDBError(err)
	}
	// TODO: Figure out how to convert this to a struct
	return result.Item["paragraphs"], nil
}
