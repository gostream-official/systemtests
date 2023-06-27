package create

type ArtistInfo struct {
	ID        string      `json:"id" bson:"_id"`
	Name      string      `json:"name" bson:"name"`
	Genres    []string    `json:"genres" bson:"genres"`
	Followers uint32      `json:"followers" bson:"followers"`
	Stats     ArtistStats `json:"stats" bson:"stats"`
}

type ArtistStats struct {
	Popularity float32 `json:"popularity" bson:"popularity"`
}

type CreateArtistInfo struct {
	Name      string      `json:"name" bson:"name"`
	Genres    []string    `json:"genres" bson:"genres"`
	Followers uint32      `json:"followers" bson:"followers"`
	Stats     ArtistStats `json:"stats" bson:"stats"`
}
