package routes

import (
	"ltv_challenge/controllers"

	"goji.io"
	"goji.io/pat"
)

// Init the songs api routes
func set_song_routes(songs *goji.Mux) {
	songs.HandleFunc(pat.Get("/artist/:artist"), controllers.SongByArtist)
	songs.HandleFunc(pat.Get("/genre/:genre"), controllers.SongByGenre)
	songs.HandleFunc(pat.Get("/name/:name"), controllers.SongByName)
	songs.HandleFunc(pat.Get("/length"), controllers.SongByLength)
}
