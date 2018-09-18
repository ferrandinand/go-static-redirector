package routers

import (
	"github.com/gorilla/mux"
)

func InitRouters(redirectionFile string) *mux.Router {

	router := mux.NewRouter().StrictSlash(false)

	router = SetRouters(router, redirectionFile)

	return router

}
