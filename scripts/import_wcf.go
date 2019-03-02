package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

func parseWCF(data []byte) []wcfChapter {

}

func parseParagraphs(data []byte, wcf []wcfChapter) []wcfChapter {
	confession := wcf
	sliceOfWords := strings.Fields(string(data))
	indexesForEachParagraph := []int{}
	for i, word := range sliceOfWords {
		if word == "__WCF_PARAGRAPH__" {
			indexesForEachParagraph = append(indexesForEachParagraph, i+1)
		}
	}
	fmt.Println("length of paragraphs: ", len(indexesForEachParagraph))
	for i, indexForParagraph := range indexesForEachParagraph {
		var paragraph []string
		if i != len(indexesForEachParagraph)-1 {
			paragraph = sliceOfWords[indexForParagraph:indexesForEachParagraph[i+1]]
		} else {
			paragraph = sliceOfWords[:indexForParagraph]
		}
		// confession = append(confession, wcfChapter{})
		// confession.find(() => )
	}
	return confession
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
	content, err := ioutil.ReadFile("WCF.txt")
	if err != nil {
		log.Fatal(err)
	}

	wcf := parseWCF(content)

	fmt.Println(string(content))
	fmt.Println(wcf)
}
