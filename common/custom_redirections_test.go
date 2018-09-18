package common

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonEquivalencies = []struct {
	Path string
	URL  string
}{
	{Path: "/health", URL: "http://www.flywire.com/healthcare"},
	{Path: "/text", URL: "/music"},
}

func TestLoadFile(t *testing.T) {

	redirections := LoadFile("routes_test.json")
	var r Router

	err := json.Unmarshal(redirections, &r)
	if err != nil {
		t.Error("Unable to load json file in the proper format")
	}

	for _, url := range jsonEquivalencies {
		for _, urlFile := range r.Redirections {
			if url.Path == urlFile.Path {
				if url.URL != urlFile.URL {
					t.Error("Expected same path for " + url.Path + " but found different urls")
					fmt.Print("For url: " + url.URL)
					fmt.Print("\nFound: " + urlFile.URL)
				}
			}

		}

	}
}
