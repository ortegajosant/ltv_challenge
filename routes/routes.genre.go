package routes

import (
	"ltv_challenge/controllers"

	"goji.io"
	"goji.io/pat"
)

func set_genres_routes(genres *goji.Mux) {
	genres.HandleFunc(pat.Get("/genres/search/:set"), controllers.Genre_search)
}
