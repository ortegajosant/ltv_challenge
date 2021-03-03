package routes

// This is the main script that manage the routes of the API
// Here the routes are setted

import (
	"goji.io"
	"goji.io/pat"
)

// Init the routes of the API
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
