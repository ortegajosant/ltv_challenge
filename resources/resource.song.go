package resources

// This is a Presentation Model Resource to Song Model
type SongsResource struct {
	Artist string `json:"artist"`
	Song   string `json:"song"`
	Genre  string `json:"genre"`
	Length int    `json:"length"`
}
