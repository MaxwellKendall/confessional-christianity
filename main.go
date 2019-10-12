package main

import (
	"github.com/MaxwellKendall/confessional-christianity/impl/wcf"
	"github.com/MaxwellKendall/confessional-christianity/server/endpoints"
	"github.com/MaxwellKendall/confessional-christianity/server/rest"
)

func main() {
	wcfSvc := wcf.NewServiceWithMiddleware()

	config := endpoints.Configuration{
		Wcf: wcfSvc,
	}

	endPoints := endpoints.GetEndpointConfig(config)

	rest.NewServer(endPoints)

}
