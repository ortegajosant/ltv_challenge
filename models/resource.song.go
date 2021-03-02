package models

type SongsResource struct {
	Id     int    `json:"id,omitempty"`
	Artist string `json:"artist"`
	Song   string `json:"song"`
	Genre  string `json:"genre"`
	Length int    `json:"length"`
}
