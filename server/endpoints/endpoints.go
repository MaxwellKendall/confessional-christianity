package endpoints

import (
	"context"

	"github.com/MaxwellKendall/confessional-christianity/impl/api"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints defines the name of each endpoint, and its type
type Endpoints struct {
	// every end point, defined as type endpoint.Endpoint
	GetWCFChapterEndpoint endpoint.Endpoint
}

// New returns all callable actions available via REST
func New(config Configuration) Endpoints {
	return Endpoints{
		GetWCFChapterEndpoint: makeGetWCFChapterEndpoint(config.Wcf),
	}
}

// makeGetWCFChapterEndpoint returns the endpoint
func makeGetWCFChapterEndpoint(svc api.WCFService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getWCFChapterRequest)
		res, err := svc.GetChapter(req.Chapter)
		if err != nil {
			// provide generic response type: errInvalidRequest
			return nil, err
		}
		return res, nil
	}
}
