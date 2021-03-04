package tests

// This script test the functions in the mapping model

import (
	"encoding/json"
	"ltv_challenge/mapping"
	"ltv_challenge/models"
	"ltv_challenge/resources"
	"testing"
)

// This will test the logic implemented in the mapping
func TestMapSongToSongResource(t *testing.T) {
	songModel := []models.Songs{}

	songModel = append(songModel, models.Songs{
		Id:     1,
		Artist: "ArtistTestName",
		Song:   "SongTestName",
		Genre:  2,
		Length: 100,
	})

	genreModel := []models.Genres{}
	genreModel = append(genreModel, models.Genres{
		Id:   2,
		Name: "GenreTestName",
	})

	got := mapping.MapSongToSongResource(songModel, genreModel)[0]
	want := resources.SongsResource{
		Artist: "ArtistTestName",
		Song:   "SongTestName",
		Genre:  "GenreTestName",
		Length: 100,
	}

	if got != want {
		t.Errorf("\ngot \n%q\n wanted \n%q\n", got, want)
	}
}

func TestToLengthRequestResource(t *testing.T) {
	lengthRequest := resources.LengthRequest{
		Max: 100,
		Min: 150,
	}
	j, _ := json.Marshal(lengthRequest)

	got, _ := mapping.ToLengthRequestResource(j)
	want := resources.LengthRequest{
		Max: 100,
		Min: 150,
	}
	if got != want {
		t.Errorf("\ngot \n%q\n wanted \n%q\n", got, want)
	}
}

func TestToGenresWithSongsInfo(t *testing.T) {
	genres := []models.Genres{}

	genres = append(genres, models.Genres{
		Id:   1,
		Name: "GenreNameTest",
	})

	songInfo := []resources.GenresAmountLengthResource{}
	songInfo = append(songInfo, resources.GenresAmountLengthResource{
		Number:      1,
		TotalLength: 100,
	})

	got, _ := mapping.ToGenresWithSongsInfo(genres, songInfo)
	want := resources.GenresAmountLengthResource{
		GenreName:   "GenreNameTest",
		Number:      1,
		TotalLength: 100,
	}

	if got[0] != want {
		t.Errorf("\ngot \n%q\n wanted \n%q\n", got[0], want)
	}

}
