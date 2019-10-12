package main

import (
	"fmt"
	// "strconv"

	citations "github.com/MaxwellKendall/confessional-christianity/loaders/citations"
	// db "github.com/MaxwellKendall/confessional-christianity/utils"
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

func main() {
	// svc := db.GetDBSession()
	wcfCitations := citations.ImportWcfCitations()

	// fmt.Println(wcfCitations)
	fmt.Println("HELLOOO", wcfCitations)
	// for _, chap := range wcf {
	// 	parsedChapter := WcfGetQuery{
	// 		ID:      "WCF_" + strconv.Itoa(chap.Number),
	// 		Chapter: chap.Number,
	// 	}
	// 	query, err := db.MakeQuery(parsedChapter)
	// 	if err != nil {
	// 		fmt.Println("**ERROR: ", err)
	// 	}
	// 	result, err := db.Get(svc, query)
	// 	db.HandleDBError(err)
	// 	fmt.Println(result)
	// }
}
