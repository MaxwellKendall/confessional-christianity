package main

import (
	"github.com/MaxwellKendall/confessional-christianity/api"
	"github.com/MaxwellKendall/confessional-christianity/server/endpoints"
	"github.com/MaxwellKendall/confessional-christianity/server/rest"
)

func main() {
	var svc api.WCFService

	config := endpoints.Configuration{
		Wcf: svc,
	}

	endPoints := endpoints.GetEndpointConfig(config)

	rest.NewServer(endPoints)

}
