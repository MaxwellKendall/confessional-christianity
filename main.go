// having an issue with .aws/config at work, fixed issue at work, now experiencing 400

package main

import (
	"github.com/MaxwellKendall/confessional-christianity/impl/wcf"
	"github.com/MaxwellKendall/confessional-christianity/utils"
)

func main() {
	// fmt.Println("Result", output)
	wcf.PrintWCF()
	utils.ListTables()
}
