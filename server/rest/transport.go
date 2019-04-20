package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MaxwellKendall/confessional-christianity/impl/api"
	"github.com/go-kit/kit/endpoint"
)

// the transfer from a go-struct to a json object
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// the transfer of json to a go-struct
func DecodeGetWCFChapterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getWCFChapterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}


