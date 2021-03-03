package controllers

import (
	"encoding/json"
	"fmt"
	"ltv_challenge/models"
	"ltv_challenge/persistence"
	"net/http"
	"strconv"

	"goji.io/pat"
)

// Common function to send the response
func sendGenreResponse(w http.ResponseWriter, genres []models.Genres) {
	j, err := json.Marshal(genres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)

}

// This send one
func GenreByID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(pat.Param(r, "id"))
	if err != nil {
		fmt.Fprintf(w, "Conversion error: %s!", err)
		return
	}

	genres, err := persistence.GetOneGenreById(id)

	if err != nil {
		http.Error(w, "There is an error: "+err.Error(), http.StatusBadGateway)
		return
	}

	sendGenreResponse(w, genres)
}
