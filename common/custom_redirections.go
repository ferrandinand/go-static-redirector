package common

import (
	"encoding/json"
	"log"
	"os"
)

type Redirection struct {
	Path string
	URL  string
}

type Router struct {
	Redirections []Redirection
}

var RedirectionList *Router
var logger *log.Logger

func CustomRedirections(routerFilePath string) {

	// use default logger
	logger = log.New(os.Stderr, "", log.LstdFlags)
	log.Println("Log to stderr")

	routerFile, err := os.Open(routerFilePath)
	if err != nil {
		log.Println(err)
		RedirectionList.Redirections = []Redirection{{Path: "/", URL: "."}}
	} else {
		json.NewDecoder(routerFile).Decode(&RedirectionList)
		routerFile.Close()
	}
	logger.Printf("Loading URLs...")
	routerStr, _ := json.MarshalIndent(RedirectionList, "", "  ")
	logger.Printf("Router: %s", routerStr)
}
