package persistence

// This is the repo that manage the access to Genres Table in the DB

import (
	"fmt"
	"ltv_challenge/models"

	_ "github.com/mattn/go-sqlite3"
)

// This function do the query to get the genre by ID
func GetOneGenreById(genreId int) ([]models.Genres, error) {
	query := fmt.Sprintf("SELECT * FROM genres WHERE ID = %d;", genreId)

	return getAllGenre(query)
}

// This function do the query to get the All the genres in the DB
func GetAllGenres() ([]models.Genres, error) {
	query := "SELECT * FROM Genres;"

	return getAllGenre(query)
}

// Common function to give the genres in the Genre Model Schema
func getAllGenre(query string) ([]models.Genres, error) {
	db_context := Db_init()
	rows, err := db_context.Query(query)
	if err != nil {
		return []models.Genres{}, err
	}

	// The resource is closed
	defer rows.Close()

	genres := []models.Genres{}
	var genre models.Genres

	for rows.Next() {
		rows.Scan(
			&genre.Id,
			&genre.Name,
		)

		genres = append(genres, genre)
	}

	return genres, nil
}
