package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// encodeResponse the transfer from a go-struct to a json object
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// decodeGetWCFChapterRequest the transfer of json to a go-struct
func decodeGetWCFChapterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	// var request endpoints.getWCFChapterRequest
	chapterRequested, _ := strconv.Atoi(vars["number"])

	fmt.Println("ChapterRequested: ", chapterRequested)

	ok := chapterRequested <= 33 && chapterRequested > 0
	if !ok {
		return nil, errors.New("Chapter requested does not exist")
	}

	return chapterRequested, nil
}
