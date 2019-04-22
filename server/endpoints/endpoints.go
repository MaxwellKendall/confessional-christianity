package endpoints

import (
	"context"
	"errors"
	"fmt"

	"github.com/MaxwellKendall/confessional-christianity/api"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints contains every end point
type Endpoints struct {
	GetWCFChapterEndpoint endpoint.Endpoint
}

// GetEndpointConfig returns all endpoints w/ gokit Endpoint
func GetEndpointConfig(config Configuration) Endpoints {
	return Endpoints{
		// 1. WCF end points:
		GetWCFChapterEndpoint: makeGetWCFChapterEndpoint(config.Wcf),
	}
}

// makeGetWCFChapterEndpoint returns the endpoint
func makeGetWCFChapterEndpoint(svc api.WCFService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(int)
		if !ok {
			return nil, errors.New("HERES AN ERROR")
		}

		res, err := svc.GetChapter(req)
		if err != nil {
			// provide generic response type: errInvalidRequest
			fmt.Println("*** error: ", string(err.Error()))
			return nil, err
		}
		return res, nil
	}
}
