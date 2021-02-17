package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Song - Models a song
type Song struct {
	ID     primitive.ObjectID `bson:"id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Artist string             `bson:"artist,omitempty"`
	URL    string             `bson:"url,omitempty"`
	Votes  int                `bson:"votes,omitempty"`
}

// Playlist - Models a list of songs to be voted on
type Playlist struct {
	ID            primitive.ObjectID   `bson:"id,omitempty"`
	Songs         []primitive.ObjectID `bson:"songs,omitempty"`
	VoteThreshold float64              `bson:"votethreshold,omitempty"`
}
