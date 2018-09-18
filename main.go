package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ferrandinand/go-static-redirector/routers"
	"github.com/urfave/negroni"
)

func main() {
	redirectionFile := os.Getenv("REDIRECTION_FILE")
	applicationPort := os.Getenv("APP_PORT")

	r := routers.InitRouters(redirectionFile)

	n := negroni.Classic()
	n.Use(negroni.NewStatic(http.Dir("/public")))
	n.UseHandler(r)

	log.Println("Listening...")
	http.ListenAndServe(":"+applicationPort, n)
}
