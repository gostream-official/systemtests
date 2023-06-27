package create

import "time"

type AlbumInfo struct {
	ID       string     `json:"id" bson:"_id"`
	Title    string     `json:"title" bson:"title"`
	TrackIDs []string   `json:"trackIds" bson:"trackIds"`
	Stats    AlbumStats `json:"stats" bson:"stats"`
}

type AlbumStats struct {
	Popularity float32 `json:"popularity" bson:"popularity,truncate"`
}

type CreateAlbumInfo struct {
	Title    string     `json:"title" bson:"title"`
	TrackIDs []string   `json:"trackIds" bson:"trackIds"`
	Stats    AlbumStats `json:"stats" bson:"stats"`
}

type TrackInfo struct {
	ID                string        `json:"id" bson:"_id"`
	ArtistID          string        `json:"artistId" bson:"artistId"`
	FeaturedArtistIDs []string      `json:"featuredArtistIds" bson:"featuredArtistIds"`
	Title             string        `json:"title" bson:"title"`
	Label             string        `json:"label" bson:"label"`
	ReleaseDate       time.Time     `json:"releaseDate" bson:"releaseDate"`
	TrackStats        TrackStats    `json:"trackStats" bson:"trackStats"`
	AudioFeatures     AudioFeatures `json:"audioFeatures" bson:"audioFeatures"`
}

type CreateTrackInfo struct {
	ArtistID          string        `json:"artistId" bson:"artistId"`
	FeaturedArtistIDs []string      `json:"featuredArtistIds" bson:"featuredArtistIds"`
	Title             string        `json:"title" bson:"title"`
	Label             string        `json:"label" bson:"label"`
	ReleaseDate       string        `json:"releaseDate" bson:"releaseDate"`
	TrackStats        TrackStats    `json:"trackStats" bson:"trackStats"`
	AudioFeatures     AudioFeatures `json:"audioFeatures" bson:"audioFeatures"`
}

type TrackStats struct {
	Streams uint32 `json:"streams" bson:"streams"`
	Likes   uint32 `json:"likes" bson:"likes"`
}

type AudioFeatures struct {
	Key              string  `json:"key" bson:"key"`
	Tempo            float32 `json:"tempo" bson:"tempo"`
	Duration         float32 `json:"duration" bson:"duration"`
	Energy           float32 `json:"energy" bson:"energy"`
	Danceability     float32 `json:"danceability" bson:"danceability"`
	Accousticness    float32 `json:"accousticness" bson:"accousticness"`
	Instrumentalness float32 `json:"instrumentalness" bson:"instrumentalness"`
	Liveness         float32 `json:"liveness" bson:"liveness"`
	Loudness         float32 `json:"loudness" bson:"loudness"`
	TimeSignature    int     `json:"timeSignature" bson:"timeSignature"`
}
