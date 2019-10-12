package rest

import (
	"log"
	"net/http"
	"time"

	gokit "github.com/go-kit/kit/transport/http"

	"github.com/gorilla/mux"

	"github.com/MaxwellKendall/confessional-christianity/server/endpoints"
)

const (
	baseURL = "/api"
	WcfURL  = "/wcf"
)

// NewServer creates a server for all this mess
func NewServer(endpoints endpoints.Endpoints) *http.Server {
	r := makeHandlers(endpoints) // create custom router using gorilla
	server := &http.Server{
		Addr:           "localhost:1517", // port address
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())

	return server
}

func makeHandlers(endpoints endpoints.Endpoints) *mux.Router {
	r := mux.NewRouter()

	getWcfChapterHandler := gokit.NewServer(
		endpoints.GetWCFChapterEndpoint,
		decodeGetWCFChapterRequest,
		encodeResponse,
	)

	r.Methods("GET").Path(baseURL + WcfURL + "/chapter/{number}").Handler(getWcfChapterHandler)

	return r
}
