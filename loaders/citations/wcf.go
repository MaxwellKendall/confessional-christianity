package citations

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/MaxwellKendall/confessional-christianity/api"
)

const (
	wcfChapterAnnotation            = "__WCF_CHAPTER__"
	wcfParagraphAnnotation          = "__WCF_PARAGRAPH__"
	wcfScriptureReferenceAnnotation = "__WCF_SCRIPTURE_REF"
	wcfScriptureProofAnnotation     = "WCF_PROOF"
)

// const FindScriptureBook = "(?P<book>((1.{1}[A-Z][a-z]*)|(2\s[A-Z][a-z]*))|[A-Z][a-z]*)"
// const FindScriptureVerses = "(?P<verse>(\d{1,3}:\d{1,3}-\d{1,3}|\d{1,3}:\d{1,3})(:\d{1,3}|(,\s\d{1,3}|\4-\d{1,3})*|\b))"
// const regexString = "(?P<citation>{book}(\.\s|\s){verse})".format(book=FindScriptureBook, verse=FindScriptureVerses)

func getReferenceIDsByParagraphNumber(chapter string) map[string][]string {
	rtrn := map[string][]string{}
	arrayOfParagraphsForChapter := strings.Split(chapter, wcfParagraphAnnotation)
	for i, paragraph := range arrayOfParagraphsForChapter[1:] {
		// skips the first item which is the title of the chapter
		paragraphNumber := strconv.Itoa(i + 1)
		arrayOfParagraphWords := strings.Fields(paragraph)
		for i, wordInParagraph := range arrayOfParagraphWords {
			if wordInParagraph == wcfScriptureReferenceAnnotation {
				rtrn[paragraphNumber] = append(rtrn[paragraphNumber], arrayOfParagraphWords[i+1])
			}
		}
	}
	return rtrn
}

func getScripturesByReferenceID(chapter string) map[string][]string {
	rtrn := map[string][]string{}
	arrayOfProofsForChapter := strings.Split(chapter, wcfScriptureProofAnnotation)
	for _, proof := range arrayOfProofsForChapter[1:] {
		// skip the first item which is the title of the chapter
		citationID := string(proof[1])
		// TODO: utilize regex to parse the book and verse data
		scripture := "Jn 3:16"
		rtrn[citationID] = append(rtrn[citationID], scripture)
	}
	return rtrn
}

func parseWcfCitationsByChapter(data []byte) []api.Citation {
	rtrn := []api.Citation{}
	arrayOfChapters := strings.Split(string(data), wcfChapterAnnotation)[1:]
	for i, chapter := range arrayOfChapters[:1] {
		chapterNumber := strconv.Itoa(i + 1)
		headingID := "WCF_" + chapterNumber
		referenceIDsByParagraphNumber := getReferenceIDsByParagraphNumber(chapter)
		scripturesByReferenceID := getScripturesByReferenceID(chapter)
		for paragraphNumber := range referenceIDsByParagraphNumber {
			for _, referenceID := range referenceIDsByParagraphNumber[paragraphNumber] {
				parsedReferenceID := string(referenceID[0]) // excludes '.'
				newCitation := api.Citation{
					ID:           headingID + "_" + paragraphNumber + "_" + string(referenceID),
					ConfessionID: "WCF",
					HeadingID:    headingID,
					PassageID:    headingID + "_" + paragraphNumber,
					ReferenceID:  parsedReferenceID,
					Scripture:    scripturesByReferenceID[parsedReferenceID],
					Tags:         []string{"baptism"},
				}
				rtrn = append(rtrn, newCitation)
			}
		}
	}

	return rtrn
}

// ImportWcfCitations returns the WCF as an array of wcfChapters
func ImportWcfCitations() []api.Chapter {
	data, err := ioutil.ReadFile("loaders/data/WCF.txt")
	if err != nil {
		log.Fatal(err)
	}

	citations := parseWcfCitationsByChapter(data)

	for _, citation := range citations {
		fmt.Println("*****", citation.Scripture)
	}

	return []api.Chapter{}
}
