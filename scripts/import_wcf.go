package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type wcfParagraph struct {
	Content         string            `json:"content"`
	ScriptureProofs map[string]string `json:"scripture_proofs"`
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

func parseTitles(data []byte) []wcfChapter {
	sliceOfWords := strings.Fields(string(data))
	begin := 0
	end := 1
	chapterNumber := 0
	confession := []wcfChapter{}
	for i, word := range sliceOfWords {
		if word == "__WCF_CHAPTER__" {
			begin = i + 1
			for x, nextWord := range sliceOfWords[begin:] {
				if nextWord == "__WCF_PARAGRAPH__" {
					chapterNumber++
					end = x + begin
					title := strings.Join(sliceOfWords[begin:end], " ")
					confession = append(confession, wcfChapter{Title: title, Number: chapterNumber})
					break
				}
			}
		}
	}
	return confession
}

func main() {
	wcf := []wcfChapter{}
	content, err := ioutil.ReadFile("WCF.txt")
	if err != nil {
		log.Fatal(err)
	}

	wcf = parseTitles(content)

	fmt.Println(string(content))
	fmt.Println(wcf)
}
