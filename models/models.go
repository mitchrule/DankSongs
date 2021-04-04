package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User - Models a user and their associated credentials
type User struct {
	ID       primitive.ObjectID `bson:"id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Password string             `bson:"password"`
}

type Vote struct {
	ID   primitive.ObjectID `bson:"id,omitempty"`
	Name string             `bson:"name,omitempty"`
	Time primitive.DateTime `bson:"time,omitempty"`
}

// Claims - for the JWT token verification
type Claims struct {
	ID       primitive.ObjectID `bson:"id,omitempty"`
	Username string             `json:"username"`
	jwt.StandardClaims
}

// Song - Models a song
type Song struct {
	ID     primitive.ObjectID `bson:"id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Artist string             `bson:"artist,omitempty"`
	URL    string             `bson:"url,omitempty"`
	Votes  []Vote             `bson:"votes,omitempty"`
}

// Playlist - Models a list of songs to be voted on
type Playlist struct {
	ID            primitive.ObjectID `bson:"id,omitempty"`
	Name          string             `bson:"string,omitempty"`
	Songs         []Song             `bson:"songs,omitempty"`
	VoteThreshold float64            `bson:"votethreshold,omitempty"`
}
