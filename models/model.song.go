package models

type Songs struct {
	Id     int    `json:"id,omitempty"`
	Artist string `json:"artist"`
	Song   string `json:"song"`
	Genre  int    `json:"genre,omitempty"`
	Length int    `json:"length"`
}
