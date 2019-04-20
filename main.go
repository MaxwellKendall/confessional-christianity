package main

import (
	"net/http"

	"github.com/MaxwellKendall/confessional-christianity/impl/wcf"
	"github.com/MaxwellKendall/confessional-christianity/server/rest"

	gktransport "github.com/go-kit/kit/transport/http"
)

func main() {
	var svc wcf.Service

	getWCFChapterHandler := gktransport.NewServer(
		rest.MakeGetWCFChapterEndpoint(svc),
		rest.DecodeGetWCFChapterRequest,
		rest.EncodeResponse,
	)

	http.Handle("/chapter", getWCFChapterHandler)
	http.ListenAndServe(":1517", nil)
}
