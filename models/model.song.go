package models

// This is the model of the Songs DB Table
type Songs struct {
	Id     int    `json:"id,omitempty"`
	Artist string `json:"artist"`
	Song   string `json:"song"`
	Genre  int    `json:"genre,omitempty"`
	Length int    `json:"length"`
}
