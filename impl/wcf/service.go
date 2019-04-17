package wcf

import (
	"fmt"

	"github.com/MaxwellKendall/confessional-christianity/impl/wcf/dao"
)

func PrintWCF() {
	wcfChapter, _ := ccdb.GetWcfChapter(1)
	fmt.Println(wcfChapter)
}
