package routes

import (
	"ltv_challenge/controllers"

	"goji.io"
	"goji.io/pat"
)

// Init the genre api routes
func set_genres_routes(genres *goji.Mux) {
	genres.HandleFunc(pat.Get("/search/:id"), controllers.GenreByID)
}
