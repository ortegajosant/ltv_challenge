package resources

// This is a Presentation Model Resource to Song Model
type GenresAmountLengthResource struct {
	GenreName   string `json:"genre,omitempty"`
	Number      int    `json:"number"`
	TotalLength int    `json:"totalLength"`
}
