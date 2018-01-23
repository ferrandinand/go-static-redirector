package main

import (
	"log"
	"net/http"
	"os"

	"../cms-middleware/common"
	"../cms-middleware/routers"
	"github.com/urfave/negroni"
)

func main() {
	redirectionFile := os.Getenv("REDIRECTION_FILE")
	applicationPort := os.Getenv("APP_PORT")

	common.CustomRedirections(redirectionFile)
	r := routers.InitRouters()

	n := negroni.Classic()
	n.Use(negroni.NewStatic(http.Dir("/public")))
	n.UseHandler(r)

	log.Println("Listening...")
	http.ListenAndServe(":"+applicationPort, n)
}
