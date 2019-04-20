package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/MaxwellKendall/confessional-christianity/impl/api"
	"github.com/go-kit/kit/endpoint"
)

// the required shape of a request to getWCFChapter
type getWCFChapterRequest struct {
	Chapter int `json:"chapter"`
}

// the shape of the response to the getWCFChapter request
type getWCFChapterResponse struct {
	V   api.WCFChapter `json:"v"`
	Err string         `json:"err,omitempty"`
}

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

// MakeGetWCFChapterEndpoint returns the endpoint
func MakeGetWCFChapterEndpoint(svc api.WCFService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getWCFChapterRequest)
		v, err := svc.GetChapter(req.Chapter)
		if err != nil {
			return getWCFChapterResponse{v, err.Error()}, nil
		}
		return getWCFChapterResponse{v, ""}, nil
	}
}
