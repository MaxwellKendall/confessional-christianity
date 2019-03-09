// TODOs:
// 1. Filter out annotations
// 2. Inculude scripture proofs
// 3. Get DB Connection
// 4. Post to DB

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

const (
	wcfChapterAnnotation            = "__WCF_CHAPTER__"
	wcfParagraphAnnotation          = "__WCF_PARAGRAPH__"
	wcfScriptureReferenceAnnotation = "__WCF_SCRIPTURE_REF"
	wcfScriptureProofAnnotation     = "WCF_PROOF"
)

func parseWCF(data []byte) []wcfChapter {
	wcfWords := strings.Fields(string(data))
	confession := []wcfChapter{}
	chapterIndexStart := 0
	chapterIndexEnd := 1
	chapterNumber := 0
	for i, word := range wcfWords {
		if word == wcfChapterAnnotation {
			chapterNumber++
			chapterIndexStart = i
			for x, nextWord := range wcfWords[chapterIndexStart+1:] {
				currentIndexRelativeToWcfWords := chapterIndexStart + x + 1
				if nextWord == wcfChapterAnnotation {
					chapterIndexEnd = x + chapterIndexStart
					break
				} else if currentIndexRelativeToWcfWords >= len(wcfWords)-1 {
					chapterIndexEnd = len(wcfWords) - 1
					break
				}
			}

			newChapter := wcfChapter{
				Title:      getChapterTitle(wcfWords[chapterIndexStart:chapterIndexEnd]),
				Number:     chapterNumber,
				Paragraphs: getChapterParagraph(wcfWords[chapterIndexStart:chapterIndexEnd]),
			}

			confession = append(confession, newChapter)
		}
	}
	return confession
}

func getScriptureProofs(wcfWords []string) map[string]string {
	rtrn := make(map[string]string)
	// fmt.Println("heyheyhey", wcfWords)
	return rtrn
}

func filterAnnotations(wcfWords []string, annotationType string) []string {
	rtrn := make([]string, 1)
	for i, word := range wcfWords {
		if i != 0 && wcfWords[i-1] == annotationType {
			alphabeticReference := strings.Split(wcfWords[i], ".")
			parsedWord := strings.Replace(wcfWords[i], wcfWords[i], "("+alphabeticReference[0]+")", -1)
			rtrn = append(rtrn, parsedWord)
		} else if word != annotationType {
			rtrn = append(rtrn, word)
		}
	}
	return rtrn
}

func getChapterParagraph(wcfWords []string) []wcfParagraph {
	paragraphs := []wcfParagraph{}
	paragraphIndexStart := 0
	paragraphIndexEnd := 1
	for i, word := range wcfWords {
		if word == wcfParagraphAnnotation {
			paragraphIndexStart = i + 1
			for x, nextWord := range wcfWords[paragraphIndexStart:] {
				if nextWord == wcfParagraphAnnotation || nextWord == wcfScriptureProofAnnotation {
					paragraphIndexEnd = x + paragraphIndexStart
					break
				}
			}
			newParagraph := wcfParagraph{
				Content:         strings.Join(filterAnnotations(wcfWords[paragraphIndexStart:paragraphIndexEnd], wcfScriptureReferenceAnnotation), " "),
				ScriptureProofs: getScriptureProofs(wcfWords[paragraphIndexEnd:]),
			}
			paragraphs = append(paragraphs, newParagraph)
		}
	}

	return paragraphs
}

func getChapterTitle(wcfWords []string) string {
	titleIndexStart := 0
	titleIndexEnd := 1

	title := ""
	for i, word := range wcfWords {
		if word == wcfChapterAnnotation {
			titleIndexStart = i + 1
			for x, nextWord := range wcfWords[titleIndexStart:] {
				if nextWord == wcfParagraphAnnotation {
					titleIndexEnd = titleIndexStart + x
					break
				}
			}
		}
		title = strings.Join(wcfWords[titleIndexStart:titleIndexEnd], " ")
	}
	return title
}

func main() {
	content, err := ioutil.ReadFile("WCF.txt")
	if err != nil {
		log.Fatal(err)
	}

	wcf := parseWCF(content)

	fmt.Println("parsed file :", wcf[0].Paragraphs[0])
}
