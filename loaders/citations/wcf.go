package citations

import (
	"io/ioutil"
	"log"
	"regexp"
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

var scriptureAbbreviationMap = map[string]string{
	"Gen":       "Genesis",
	"Ex":        "Exodus",
	"Lev":       "Leviticus",
	"Num":       "Numbers",
	"Deut":      "Deuteronomy",
	"Josh":      "Joshua",
	"Jdg":       "Judges",
	"Rth":       "Ruth",
	"1 Sam":     "1 Samuel",
	"2 Sam":     "2 Samuel",
	"1 Kings":   "1 Kings",
	"2 Kings":   "2 Kings",
	"1 Chron":   "1 Chronicles",
	"2 Chron":   "2 Chronicles",
	"Ezra":      "Ezra",
	"Neh":       "Nehemiah",
	"Est":       "Esther",
	"Job":       "Job",
	"Ps":        "Psalms",
	"Prov":      "Proverbs",
	"Eccl":      "Ecclesiastes",
	"SoS":       "Song of Solomon",
	"Isa":       "Isaiah",
	"Jer":       "Jeremiah",
	"Lam":       "Lamentations",
	"Ezek":      "Ezekiel",
	"Dan":       "Daniel",
	"Hos":       "Hosea",
	"Joel":      "Joel",
	"Amos":      "Amos",
	"Obadiah":   "Obadiah",
	"Jonah":     "Jonah",
	"Mic":       "Micah",
	"Nah":       "Nahum",
	"Hab":       "Habakkuk",
	"Zephaniah": "Zephaniah",
	"Hag":       "Haggai",
	"Zech":      "Zechariah",
	"Mal":       "Malachi",
	"Matt":      "Matthew",
	"Mark":      "Mark",
	"Luke":      "Luke",
	"John":      "John",
	"Acts":      "Acts",
	"Rom":       "Romans",
	"1 Cor":     "1 Corinthians",
	"2 Cor":     "2 Corinthians",
	"Gal":       "Galatians",
	"Eph":       "Ephesians",
	"Phil":      "Philippians",
	"Col":       "Colossians",
	"1 Thess":   "1 Thessalonians",
	"2 Thess":   "2 Thessalonians",
	"1 Tim":     "1 Timothy",
	"2 Tim":     "2 Timothy",
	"Titus":     "Titus",
	"Philemon":  "Philemon",
	"Heb":       "Hebrews",
	"James":     "James",
	"1 Pet":     "1 Peter",
	"2 Pet":     "2 Peter",
	"1 John":    "1 John",
	"2 John":    "2 John",
	"3 John":    "3 John",
	"Jude":      "Jude",
	"Rev":       "Revelation",
}

var regexString = regexp.MustCompile(`(?P<book>((1.{1}[A-Z][a-z]*)|(2\s[A-Z][a-z]*)|3\s[A-Z][a-z]*)|[A-Z][a-z]*)(\.\s|\s)(?P<verse>(\d{1,3}:\d{1,3}-\d{1,3}|\d{1,3}:\d{1,3})(:\d{1,3}|(,\s\d{1,3}|4-\d{1,3})*|\b))`)

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
	sliceOfProofsInChapter := strings.Split(chapter, wcfScriptureProofAnnotation)
	for _, proof := range sliceOfProofsInChapter[1:] {
		// skip the first item which is the title of the chapter
		citationID := string(proof[1])
		citedScripture := []string{}
		matchedItems := regexString.FindAllStringSubmatch(proof, -1)
		bookIndex := 0
		verseIndex := 0
		for i, name := range regexString.SubexpNames() {
			if i == 0 || name == "" {
				continue
			}
			if name == "book" {
				bookIndex = i
			}
			if name == "verse" {
				verseIndex = i
			}
		}

		for _, match := range matchedItems {
			abbreviatedBook := match[bookIndex]
			fullBook := scriptureAbbreviationMap[abbreviatedBook]
			citedScripture = append(citedScripture, fullBook+" "+match[verseIndex])
		}
		rtrn[citationID] = citedScripture
	}
	return rtrn
}

func parseWcfCitationsByChapter(data []byte) []api.Citation {
	rtrn := []api.Citation{}
	arrayOfChapters := strings.Split(string(data), wcfChapterAnnotation)[1:]
	for i, chapter := range arrayOfChapters {
		chapterNumber := strconv.Itoa(i + 1)
		headingID := "WCF_" + chapterNumber
		referenceIDsByParagraphNumber := getReferenceIDsByParagraphNumber(chapter)
		scripturesByReferenceID := getScripturesByReferenceID(chapter)
		for paragraphNumber := range referenceIDsByParagraphNumber {
			for _, referenceID := range referenceIDsByParagraphNumber[paragraphNumber] {
				parsedReferenceID := string(referenceID[0]) // excludes '.'
				headingNumber, _ := strconv.ParseInt(chapterNumber, 10, 64)
				paragraphNumberInt, _ := strconv.ParseInt(paragraphNumber, 10, 64)
				newCitation := api.Citation{
					ID:            headingID + "_" + paragraphNumber + "_" + string(referenceID),
					ConfessionID:  "WCF",
				HeadingID:     headingID,
					PassageID:     headingID + "_" + paragraphNumber,
					HeadingNumber: headingNumber,
					PassageNumber: paragraphNumberInt,
					ReferenceID:   parsedReferenceID,
					Scripture:     scripturesByReferenceID[parsedReferenceID],
					Tags:          []string{"baptism"},
				}
				rtrn = append(rtrn, newCitation)
			}
		}
	}

	return rtrn
}

// ImportWcfCitations returns the WCF as an array of wcfChapters
func ImportWcfCitations() []api.Citation {
	data, err := ioutil.ReadFile("loaders/data/WCF.txt")
	if err != nil {
		log.Fatal(err)
	}

	return parseWcfCitationsByChapter(data)
}
