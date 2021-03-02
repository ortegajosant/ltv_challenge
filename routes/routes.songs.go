package routes

import (
	"ltv_challenge/controllers"

	"goji.io"
	"goji.io/pat"
)

func set_song_routes(songs *goji.Mux) {
	songs.HandleFunc(pat.Get("/get/artist/:artist"), controllers.SongByArtist)
	songs.HandleFunc(pat.Get("/get/genre/:genre"), controllers.SongByGenre)
	songs.HandleFunc(pat.Get("/get/name/:name"), controllers.SongByName)
}
