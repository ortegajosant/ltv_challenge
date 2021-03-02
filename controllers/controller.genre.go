package controllers

import (
	"fmt"
	"net/http"

	"goji.io/pat"
)

func Genre_search(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}
