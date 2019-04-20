package wcf

import (
	"encoding/json"
	"fmt"

	"github.com/MaxwellKendall/confessional-christianity/impl/api"

	ccdb "github.com/MaxwellKendall/confessional-christianity/impl/wcf/dao"
)

// The Service is used in the mux
type Service struct{}

// GetChapter returns a chapter of the WCF
func (Service) GetChapter(chapter int) (api.WCFChapter, error) {
	wcfChapter, _ := ccdb.GetWcfChapter(chapter)
	json, err := json.Marshal(wcfChapter)
	if err != nil {
		return api.WCFChapter{}, err
	}
	fmt.Println(string(json))
	return wcfChapter, nil
}
