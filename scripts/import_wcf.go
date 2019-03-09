// TODOs:
// 1. Scripture proofs are currently modled to be mapped by paragraph, but should be mapped by chapter?
// 2. Get DB Connection
// 3. Post to DB

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	wcfChapterAnnotation            = "__WCF_CHAPTER__"
	wcfParagraphAnnotation          = "__WCF_PARAGRAPH__"
	wcfScriptureReferenceAnnotation = "__WCF_SCRIPTURE_REF"
	wcfScriptureProofAnnotation     = "WCF_PROOF"
)

type wcfChapter struct {
	Title      string            `json:"title"`
	Number     int               `json:"number"`
	Paragraphs []string          `json:"paragraphs"`
	Proofs     map[string]string `json:"proofs"`
}

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
				Proofs:     getScriptureProofs(wcfWords[chapterIndexStart:chapterIndexEnd]),
			}

			confession = append(confession, newChapter)
		}
	}
	return confession
}

func getScriptureProofs(wcfWords []string) map[string]string {
	rtrn := make(map[string]string)
	for i, word := range wcfWords {
		if word == wcfScriptureProofAnnotation {
			counter := 0
			startIndex := i + 2 // not including alphabetical reference id as it will be the key in the map itself
			endIndex := 0
			key := strings.Split(wcfWords[i+1], ".")[0]

			for x, nextWord := range wcfWords[startIndex:] {
				if x == len(wcfWords[startIndex:])-1 {
					endIndex = startIndex + x
					break
				} else if nextWord == wcfScriptureProofAnnotation {
					endIndex = startIndex + x
					break
				}
			}
			filteredProofs := filterAnnotations(wcfWords[startIndex:endIndex], wcfScriptureProofAnnotation)
			rtrn[key] = strings.Join(filteredProofs, " ")
			counter++
		}
	}
	return rtrn
}

func filterAnnotations(wcfWords []string, annotationType string) []string {
	rtrn := make([]string, 1)
	for i, word := range wcfWords {
		switch annotationType {
		case wcfScriptureReferenceAnnotation:
			if i != 0 && wcfWords[i-1] == annotationType {
				alphabeticReference := strings.Split(wcfWords[i], ".")
				parsedWord := strings.Replace(wcfWords[i], wcfWords[i], "("+alphabeticReference[0]+")", -1)
				rtrn = append(rtrn, parsedWord)
			} else if word != annotationType {
				rtrn = append(rtrn, word)
			}
		case wcfScriptureProofAnnotation:
			if word != annotationType {
				rtrn = append(rtrn, word)
			}
		}
	}
	return rtrn
}

func getChapterParagraph(wcfWords []string) []string {
	paragraphs := []string{}
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
			newParagraph := strings.Join(filterAnnotations(wcfWords[paragraphIndexStart:paragraphIndexEnd], wcfScriptureReferenceAnnotation), " ")
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
	fmt.Println("*************************************** Title:", wcf[13].Title)
	fmt.Println("*************************************** Content:", wcf[13].Paragraphs[0])
	fmt.Println("*************************************** Proof A:", wcf[13].Proofs["a"])
	fmt.Println("*************************************** Proof B:", wcf[13].Proofs["b"])
	fmt.Println("*************************************** Proof C:", wcf[13].Proofs["c"])
	fmt.Println("*************************************** Proof D:", wcf[13].Proofs["d"])
	fmt.Println("*************************************** Number of Proofs:", len(wcf[13].Proofs))

	test, err := json.Marshal(wcf[0].Proofs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(test))
}
