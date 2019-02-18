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

func parseParagraphs(data []byte) []byte {
	sliceOfWords := strings.Fields(string(data))
	var token []byte
	indexesForEachParagraph := []int{}
	for i, word := range sliceOfWords {
		if word == "__WCF_PARAGRAPH__" {
			indexesForEachParagraph = append(indexesForEachParagraph, i+1)
		}
	}
	fmt.Println("length of paragraphs: ", len(indexesForEachParagraph))
	for i, indexForParagraph := range indexesForEachParagraph {
		paragraph := sliceOfWords[indexForParagraph:indexesForEachParagraph[i+1]]
		fmt.Println("YO ", paragraph)
	}
	return token
}

func parseTitles(data []byte) []byte {
	var token []byte
	sliceOfWords := strings.Fields(string(data))
	begin := 0
	end := 1
	for i, word := range sliceOfWords {
		if word == "__WCF_CHAPTER__" {
			begin = i + 1
			for x, nextWord := range sliceOfWords[begin:] {
				if nextWord == "__WCF_PARAGRAPH__" || strings.HasPrefix("__WCF", nextWord) {
					end = x + begin
					title := []byte(strings.Join(sliceOfWords[begin:end], " "))
					token = append(token, title...)
					break
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

	wcfToken := parseTitles(data)
	// wcfToken = append(wcfToken, parseParagraphs(data)...)
	// arrayOfScriptureReferences :=

	return len(data), wcfToken, nil
}

func main() {
	// perhaps should not use scan here. Ran into an issue where split fn was not working only because the data chunk contained only part of a word
	f, _ := os.Open("./WCF.txt")
	s := bufio.NewScanner(f)
	s.Split(splitWCF)
	wcf := []wcfChapter{}
	for s.Scan() {
		// Right now just getting the chapter title, but would ideally like to grab the entire chapter and return it as a text block
		wcf = append(wcf, wcfChapter{Title: s.Text()})
	}
	fmt.Println(wcf, "length ", len(wcf))
}
