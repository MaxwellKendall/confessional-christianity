package endpoints

// all request types go here

// the required shape of a request to getWCFChapter
type getWCFChapterRequest struct {
	Chapter int `json:"chapter"`
}
