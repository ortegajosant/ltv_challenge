package persistence

// This is the repo that manage the access to Songs Table in the DB

import (
	"fmt"
	"ltv_challenge/models"

	_ "github.com/mattn/go-sqlite3"
)

// This function do the query to get the songs by artist
func GetSongByArtist(artist string) ([]models.Songs, error) {
	query := fmt.Sprintf("SELECT * FROM songs WHERE artist = '%s';", artist)

	return getAllSongs(query)
}

// This function do the query to get the songs by genre
func GetSongByGenre(genreName string) ([]models.Songs, error) {
	query := fmt.Sprintf("SELECT S.ID, S.artist, S.song, S.genre, S.length FROM songs S INNER JOIN genres G on G.ID = S.genre WHERE G.name = '%s';", genreName)
	return getAllSongs(query)
}

// This function do the query to get the songs by song name
func GetSongByName(songName string) ([]models.Songs, error) {
	query := fmt.Sprintf("SELECT * FROM songs WHERE song = '%s';", songName)

	return getAllSongs(query)
}

// Common function to give the genres in the Songs Model Schema
func getAllSongs(query string) ([]models.Songs, error) {
	db_context := Db_init()
	rows, err := db_context.Query(query)
	if err != nil {
		return []models.Songs{}, err
	}
	// The resource is closed
	defer rows.Close()

	songs := []models.Songs{}
	var song models.Songs

	for rows.Next() {
		rows.Scan(
			&song.Id,
			&song.Artist,
			&song.Song,
			&song.Genre,
			&song.Length,
		)

		songs = append(songs, song)
	}

	return songs, nil
}
