package mapping

// This is a quick implementation of the Presentation Model pattern
// which represent the state and behavior of the data independently of the interface

import (
	"ltv_challenge/models"
	"ltv_challenge/resources"
)

// Maps the Song Model (As is in the DB) to Song Resource (As is required)
func MapSongToSongResource(songs []models.Songs, genres []models.Genres) []resources.SongsResource {

	songResources := []resources.SongsResource{}

	var i int
	for i = 0; i < len(songs); i++ {
		songRes := resources.SongsResource{
			Artist: songs[i].Artist,
			Song:   songs[i].Song,
			Genre:  genres[i].Name,
			Length: songs[i].Length,
		}

		songResources = append(songResources, songRes)
	}

	return songResources

}
