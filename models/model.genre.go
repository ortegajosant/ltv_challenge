package models

// This is the model of the Genres DB Table
type Genres struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
}
