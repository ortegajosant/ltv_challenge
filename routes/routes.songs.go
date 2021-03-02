package routes

import (
	"ltv_challenge/controllers"

	"goji.io"
	"goji.io/pat"
)

func set_song_routes(songs *goji.Mux) {
	songs.HandleFunc(pat.Get("/search/:set"), controllers.Song_search)
}
