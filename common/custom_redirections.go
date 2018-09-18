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

//Load file gets a path for a redirection file and returns a json
func LoadFile(routerFilePath string) []byte {

	// use default logger
	logger = log.New(os.Stderr, "", log.LstdFlags)
	log.Println("Log to stderr")

	routerFile, err := os.Open(routerFilePath)
	defer routerFile.Close()

	if err != nil {
		log.Fatal(err)
	} else {
		json.NewDecoder(routerFile).Decode(&RedirectionList)
	}
	logger.Printf("Loading URLs...")

	routerStr, _ := json.MarshalIndent(&RedirectionList, "", "  ")
	logger.Printf("Router: %s", routerStr)

	return routerStr
}
