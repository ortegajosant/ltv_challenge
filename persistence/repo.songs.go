package persistence

import (
	"fmt"
	"ltv_challenge/models"

	_ "github.com/mattn/go-sqlite3"
)

func GetOneByArtist(artist string) {
	query := fmt.Sprintf("SELECT * FROM Songs WHERE artist = %s;", artist)
	getOne(query)
}

func GetOneByGenre() {

}

func GetOneBySong() {

}

func getOne(query string) (models.Songs, error) {
	rows, err := db_context.Query(query)

	if err != nil {
		return models.Songs{}, err
	}
	// Cerramos el recurso
	defer rows.Close()

	song := &models.Songs{}

	rows.Scan(
		song.Id,
		song.Artist,
		song.Song,
		song.Genre,
		song.Length,
	)

	return *song, nil
}
