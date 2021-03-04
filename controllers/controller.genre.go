package controllers

import (
	"encoding/json"
	"fmt"
	"ltv_challenge/mapping"
	"ltv_challenge/persistence"
	"net/http"
	"strconv"

	"goji.io/pat"
)

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

	json, err := json.Marshal(genres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendResponse(w, json)
}

func GenreAllInfo(w http.ResponseWriter, r *http.Request) {
	genres, err := persistence.GetAllGenres()

	if err != nil {
		http.Error(w, "There is an error: "+err.Error(), http.StatusBadGateway)
	}

	songInfo, err := getSongsInfoFromGenres(genres)

	if err != nil {
		http.Error(w, "There is an error: "+err.Error(), http.StatusBadGateway)
	}

	genresData, _ := mapping.ToGenresWithSongsInfo(genres, songInfo)

	json, err := json.Marshal(genresData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendResponse(w, json)
}
