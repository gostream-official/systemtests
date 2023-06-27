package get

type AlbumInfo struct {
	ID       string     `json:"id" bson:"_id"`
	Title    string     `json:"title" bson:"title"`
	TrackIDs []string   `json:"trackIds" bson:"trackIds"`
	Stats    AlbumStats `json:"stats" bson:"stats"`
}

type AlbumStats struct {
	Popularity float32 `json:"popularity" bson:"popularity,truncate"`
}
