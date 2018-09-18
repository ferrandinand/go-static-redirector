package routers

import (
	"encoding/json"
	"log"

	"github.com/ferrandinand/go-static-redirector/common"
	"github.com/ferrandinand/go-static-redirector/controllers"
	"github.com/gorilla/mux"
)

func SetRouters(router *mux.Router, redirectionFile string) *mux.Router {

	redirections := common.LoadFile(redirectionFile)
	var r common.Router

	err := json.Unmarshal(redirections, &r)
	if err != nil {
		log.Fatal("Unable to load json file")
	}

	for _, urlRedirection := range r.Redirections {
		router.HandleFunc(urlRedirection.Path, controllers.CustomRedirections)
	}

	return router
}
