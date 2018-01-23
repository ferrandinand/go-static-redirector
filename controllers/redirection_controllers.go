package controllers

import (
	"net/http"

	"../common"
)

//CustomRedirections request
func CustomRedirections(w http.ResponseWriter, r *http.Request) {
	for _, urlRedirection := range common.RedirectionList.Redirections {
		if r.RequestURI == urlRedirection.Path {
			http.Redirect(w, r, urlRedirection.URL, http.StatusPermanentRedirect)
		}
	}
}
