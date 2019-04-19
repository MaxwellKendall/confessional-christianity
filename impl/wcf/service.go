package wcf

import (
	"encoding/json"
	"fmt"

	"github.com/MaxwellKendall/confessional-christianity/impl/wcf/dao"
)

func PrintWCF() {
	wcfChapter, _ := ccdb.GetWcfChapter(1)
	json, err := json.Marshal(wcfChapter)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(json))
}
