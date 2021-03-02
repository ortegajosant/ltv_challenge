package controllers

import (
	"fmt"
	"net/http"

	"goji.io/pat"
)

func Song_search(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "set")
	fmt.Fprintf(w, "Hello, %s!", name)
}
