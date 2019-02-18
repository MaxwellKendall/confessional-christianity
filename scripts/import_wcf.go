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

// func parseParagraphs(data []byte) []byte {

// }

func parseTitles(data []byte) []byte {
	sliceOfWords := strings.Fields(string(data))
	endIndexForTitle := 1
	var token []byte
	for i, word := range sliceOfWords {
		// fmt.Println(len(sliceOfWords))
		if strings.HasPrefix(word, "__WCF_CHAPTER__") {
			arrayWithTitle := sliceOfWords[i+1 : i+7]
			for index, w := range arrayWithTitle {
				if w == "__WCF_PARAGRAPH__" {
					endIndexForTitle = index
					// creating a slice of bytes
					wcfHeading := []byte(strings.Join(arrayWithTitle[0:endIndexForTitle], " "))
					token = append(token, wcfHeading...)
				}
			}
		}
	}
	return token
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

	arrayOfTitles := parseTitles(data)
	// arrayOfParagraphs :=
	// arrayOfScriptureReferences :=

	return len(data), arrayOfTitles, nil
}

func main() {
	// readCSV()
	f, _ := os.Open("./WCF.txt")
	s := bufio.NewScanner(f)
	s.Split(splitWCF)
	wcf := []wcfChapter{}
	for s.Scan() {
		// Right now just getting the chapter title, but would ideally like to grab the entire chapter and return it as a text block
		fmt.Println("YO", s.Text())
		wcf = append(wcf, wcfChapter{Title: s.Text()})
	}
	fmt.Println(wcf)
}
