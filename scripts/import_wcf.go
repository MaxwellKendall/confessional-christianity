package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type wcfParagraph struct {
	Content         string            `json:"content"`
	ScriptureProofs map[string]string `json:"scripture_proofes"`
}

type wcfChapter struct {
	Title      string         `json:"title"`
	Number     int            `json:"number"`
	Paragraphs []wcfParagraph `json:"paragraphs"`
}

func readCSV() {
	// gives me a pointer to a File which implements the Reader Interface
	fileToBeRead, err := os.Open("../WCF.csv")
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

func parseTitles(data []byte, confession []wcfChapter) []wcfChapter {
	sliceOfWords := strings.Fields(string(data))
	endIndexForTitle := 1
	var wcfHeading string

	for i, word := range sliceOfWords {
		// fmt.Println(len(sliceOfWords))
		if strings.HasPrefix(word, "__WCF_CHAPTER__") {
			arrayWithTitle := sliceOfWords[i+1 : i+7]
			for index, w := range arrayWithTitle {
				if w == "1." {
					endIndexForTitle = index
					wcfHeading = strings.Join(arrayWithTitle[0:endIndexForTitle], " ")
					newChapter := wcfChapter{Title: wcfHeading}
					append(confession, newChapter)
					fmt.Println("confession", confession)
				}
			}
		}
	}
	return confession
}

func splitWCF(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// "tokenizes" the wcf into bite size pieces so we can parse it into a go struct
	if atEOF {
		fmt.Println("HEY")
		return 0, nil, nil
	}
	if err != nil {
		fmt.Println(err)
		return 0, nil, err
	}
	wcf := []wcfChapter{}

	parseTitles(data, wcf)

	return len(data), nil, nil
}

func main() {
	// readCSV()
	f, _ := os.Open("./WCF.txt")
	s := bufio.NewScanner(f)
	s.Split(splitWCF)
	s.Scan()
}
