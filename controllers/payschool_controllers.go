package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

const baseUrl = "https://www.flywire.com"

//Rewriteschool school to pay
func Rewriteschool(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["schoolId"]

	http.Redirect(w, r, baseUrl+"/pay/"+id, http.StatusPermanentRedirect)

}
