package routes

import (
	"goji.io"
	"goji.io/pat"
)

func Server_init() *goji.Mux {

	// Set the root route
	root := goji.NewMux()
	// Establish the routes for each elemento of the API
	songs := goji.SubMux()
	genres := goji.SubMux()

	// Set the subroutes
	root.Handle(pat.New("/api/songs/*"), songs)
	root.Handle(pat.New("/api/genres/*"), genres)

	// Set the routes for each element
	set_song_routes(songs)
	set_genres_routes(genres)

	return root
}
