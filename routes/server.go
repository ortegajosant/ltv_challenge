package routes

import (
	"goji.io"
	"goji.io/pat"
)

func Server_init() *goji.Mux {
	root := goji.NewMux()
	songs := goji.SubMux()
	genres := goji.SubMux()

	root.Handle(pat.New("/api/songs/*"), songs)
	root.Handle(pat.New("/api/genres/*"), genres)

	set_song_routes(songs)
	set_genres_routes(genres)

	return root
}
