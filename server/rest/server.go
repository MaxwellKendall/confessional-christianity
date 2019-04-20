package rest

import (
	"net/http"

	"github.com/MaxwellKendall/confessional-christianity/server/endpoints"
)

func NewServer(endpoints endpoints.Endpoints) *http.Server {
	r := makeHandlers(endpoints)
}

func makeHandlers(endpoints endpoints.Endpoints) *mux.Router {

}
