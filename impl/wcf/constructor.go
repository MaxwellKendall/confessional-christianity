package wcf

import (
	"github.com/MaxwellKendall/confessional-christianity/api"
)

// NewServiceWithMiddleware will eventually take a config paramater
// so that all services can have shared properties and methods
func NewServiceWithMiddleware() api.WCFService {
	var service api.WCFService
	// TODO add new properties etc via service = {}

	// adding receiver fns to satisfy the WCFService interface
	service = newService()

	return service
}
