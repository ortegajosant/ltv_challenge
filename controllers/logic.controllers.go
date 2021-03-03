package controllers

import (
	"ltv_challenge/models"
	"ltv_challenge/persistence"
)

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
