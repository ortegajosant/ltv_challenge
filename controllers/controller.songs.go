package controllers

import (
	"encoding/json"
	"fmt"
	"ltv_challenge/models"
	"ltv_challenge/persistence"
	"net/http"

	"goji.io/pat"
)

func sendResponse(w http.ResponseWriter, songs []models.Songs) {
	j, err := json.Marshal(songs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func SongByArtist(w http.ResponseWriter, r *http.Request) {
	artist := pat.Param(r, "artist")
	songs, err := persistence.GetSongByArtist(artist)

	if err != nil {
		fmt.Fprintf(w, "There is an error, %s!", err)
		return
	}

	sendResponse(w, songs)

}

func SongByGenre(w http.ResponseWriter, r *http.Request) {
	artist := pat.Param(r, "genre")
	songs, err := persistence.GetSongByGenre(artist)

	if err != nil {
		fmt.Fprintf(w, "There is an error, %s!", err)
		return
	}

	sendResponse(w, songs)

}

func SongByName(w http.ResponseWriter, r *http.Request) {
	artist := pat.Param(r, "name")
	songs, err := persistence.GetSongByName(artist)

	if err != nil {
		fmt.Fprintf(w, "There is an error, %s!", err)
		return
	}

	sendResponse(w, songs)

}
