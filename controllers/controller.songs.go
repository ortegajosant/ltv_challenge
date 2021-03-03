package controllers

import (
	"encoding/json"
	"io/ioutil"
	"ltv_challenge/mapping"
	"ltv_challenge/models"
	"ltv_challenge/persistence"
	"net/http"

	"goji.io/pat"
)

// This is a common function that helps the controller to send the responses
func sendSongResponse(w http.ResponseWriter, songs []models.Songs) {

	if len(songs) == 0 {
		http.Error(w, "Records not found", http.StatusNotFound)
		return
	}

	// Setting the song resource with presentation model quick implementation
	genres, err := getAllGenresFromSongs(songs)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resources := mapping.MapSongToSongResource(songs, genres)

	// Parse object to JSON object
	j, err := json.Marshal(resources)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Sending the file
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// The following functions get the request and return the songs
// Return the songs by artist
func SongByArtist(w http.ResponseWriter, r *http.Request) {
	artist := pat.Param(r, "artist")
	songs, err := persistence.GetSongByArtist(artist)

	if err != nil {
		http.Error(w, "There is an error: "+err.Error(), http.StatusBadGateway)
		return
	}

	sendSongResponse(w, songs)

}

// Return the songs by genre
func SongByGenre(w http.ResponseWriter, r *http.Request) {
	genre := pat.Param(r, "genre")
	songs, err := persistence.GetSongByGenre(genre)

	if err != nil {
		http.Error(w, "There is an error: "+err.Error(), http.StatusBadGateway)
		return
	}

	sendSongResponse(w, songs)

}

// Return the songs by song name
func SongByName(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	songs, err := persistence.GetSongByName(name)

	if err != nil {
		http.Error(w, "There is an error: "+err.Error(), http.StatusBadGateway)
		return
	}

	sendSongResponse(w, songs)

}

// Return the songs by song min and max length
func SongByLength(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "There is an error: "+err.Error(), http.StatusBadGateway)
		return
	}

	request, err := mapping.ToLengthRequestResource(body)

	if err != nil {
		http.Error(w, "There is an error: "+err.Error(), http.StatusBadGateway)
		return
	}

	songs, err := persistence.GetSongByLength(request.Min, request.Max)

	if err != nil {
		http.Error(w, "There is an error: "+err.Error(), http.StatusBadGateway)
		return
	}

	sendSongResponse(w, songs)

}
