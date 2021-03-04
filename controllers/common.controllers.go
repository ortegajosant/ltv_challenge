package controllers

import (
	"ltv_challenge/models"
	"ltv_challenge/persistence"
	"ltv_challenge/resources"
	"net/http"
)

// Common function to send the response
func sendResponse(w http.ResponseWriter, infoToSend []byte) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(infoToSend)

}

// This is a common function to get the genres from a list of songs
func getAllGenresFromSongs(songs []models.Songs) ([]models.Genres, error) {

	var genres []models.Genres

	for i := 0; i < len(songs); i++ {
		genre, err := persistence.GetOneGenreById(songs[i].Genre)
		if err != nil {
			return genres, err
		}

		genres = append(genres, genre[0])
	}

	return genres, nil

}

// This function get the songs Info required for each genre
func getSongsInfoFromGenres(genres []models.Genres) ([]resources.GenresAmountLengthResource, error) {

	songsInfo := []resources.GenresAmountLengthResource{}
	for i := 0; i < len(genres); i++ {
		info, err := persistence.GetSongNumberAndLength(genres[i].Name)

		if err != nil {
			return songsInfo, err
		}

		songsInfo = append(songsInfo, info[0])
	}

	return songsInfo, nil

}
