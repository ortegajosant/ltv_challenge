package persistence

import (
	"fmt"
	"ltv_challenge/models"

	_ "github.com/mattn/go-sqlite3"
)

func GetSongByArtist(artist string) ([]models.Songs, error) {
	query := fmt.Sprintf("SELECT * FROM songs WHERE artist = '%s';", artist)

	return getAll(query)
}

func GetSongByGenre(genreName string) ([]models.Songs, error) {
	query := fmt.Sprintf("SELECT S.ID, S.artist, S.song, S.genre, S.length FROM songs S INNER JOIN genres G on G.ID = S.genre WHERE G.name = '%s';", genreName)
	return getAll(query)
}

func GetSongByName(songName string) ([]models.Songs, error) {
	query := fmt.Sprintf("SELECT * FROM songs WHERE song = '%s';", songName)

	return getAll(query)
}

func getAll(query string) ([]models.Songs, error) {
	db_context := Db_init()
	rows, err := db_context.Query(query)
	if err != nil {
		return []models.Songs{}, err
	}
	// Cerramos el recurso
	defer rows.Close()

	fmt.Printf("The query is: %s\n", query)

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
