package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Might not be required if we can create the vote inside the
// song
// VotesCollection in MongoDB
//var VotesCollection *mongo.Collection = new(mongo.Collection)

// PlaylistCollection in MongoDB
var PlaylistCollection *mongo.Collection = new(mongo.Collection)

// SongsCollection in MongoDB
var SongsCollection *mongo.Collection = new(mongo.Collection)

// UsersCollection in MongoDB
var UsersCollection *mongo.Collection = new(mongo.Collection)

// UsersCollection in MongoDB
var JWTCollection *mongo.Collection = new(mongo.Collection)

// InitDatabase initialises a global database client
func InitDatabase() {
	mongoUsername := os.Getenv("MONGOUSERNAME")
	databaseName := os.Getenv("DATABASENAME")
	mongoPassword := os.Getenv("MONGOPWD")

	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.sn8oj.mongodb.net/%s?retryWrites=true&w=majority", mongoUsername, mongoPassword, databaseName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)

	// Where the collections are initalised
	PlaylistCollection = client.Database(databaseName).Collection("playlists")
	SongsCollection = client.Database(databaseName).Collection("songs")
	UsersCollection = client.Database(databaseName).Collection("users")
	JWTCollection = client.Database(databaseName).Collection("tokens")

	// log.Println("Database initialised", songsCollection)
}
