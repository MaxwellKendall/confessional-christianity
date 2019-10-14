package main

import (
	"fmt"

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
	for _, citation := range wcfCitations {
		fmt.Println(citation.Scripture)
		// TODO: Build struct that will be used to update db
		// db.MakeQuery()
		// db.Update("citations", svc)
	}
}
