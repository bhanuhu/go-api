package main

import (
	"context"
	"fmt"
	"time"

	"../../abc.ev"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Podcast struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(abc.Uri))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("quickstart")
	podcastsCollection := database.Collection("podcasts")
	podcast := Podcast{
		Title:  "The Polyglot Developer",
		Author: "Nic Raboy",
		Tags:   []string{"development", "programming", "coding"},
	}
	insertResult, err := podcastsCollection.InsertOne(ctx, podcast)
	if err != nil {
		panic(err)
	}
	fmt.Println(insertResult.InsertedID)
}
