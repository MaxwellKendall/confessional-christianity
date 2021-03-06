package wcf

import (
	"fmt"

	"github.com/MaxwellKendall/confessional-christianity/api"
	ccdb "github.com/MaxwellKendall/confessional-christianity/impl/wcf/dao"
)

// The Service is used in the mux
type service struct{}

// Exporting the service w/ this in the constructor.go
func newService() service {
	// add nifty stuff here like return service{ niftyStuff: someCoolToolToBeUsedWithinService }
	return service{}
}

// GetChapter returns a chapter of the WCF
func (service) GetChapter(chapter int) (api.WCFChapter, error) {
	wcfChapter, err := ccdb.GetWcfChapter(chapter)
	if err != nil {
		fmt.Println("Error in service layer, get Chapter", err)
		return api.WCFChapter{}, err
	}
	return wcfChapter, nil
}
