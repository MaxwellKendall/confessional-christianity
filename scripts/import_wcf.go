package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// gives me a pointer to a File which implements the Reader Interface
	fileToBeRead, err := os.Open("./WCF.csv")
	if err != nil {
		log.Fatalln(err)
	}

	csvReaderFn := csv.NewReader(fileToBeRead)
	records, err := csvReaderFn.ReadAll()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(records[0][1])
	fmt.Println("HERE")
}
