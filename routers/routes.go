package routers

import (
	"../common"
	"../controllers"
	"github.com/gorilla/mux"
)

func SetRouters(router *mux.Router) *mux.Router {

	router.HandleFunc("/school/{schoolId}", controllers.Rewriteschool)

	for _, urlRedirection := range common.RedirectionList.Redirections {
		router.HandleFunc(urlRedirection.Path, controllers.CustomRedirections)
	}

	return router
}
