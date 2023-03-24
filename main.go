package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://admin:1@coursecompose.kmyjwju.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	log.Println("Connect succesfully")

	collection := client.Database("CourseCompose").Collection("Courses")
	cur, err := collection.Find(context.Background(), bson.M{}, nil)
	if err != nil {
		log.Fatal(err)
	}

	result := []Course{}
	if err := cur.All(context.Background(), &result); err != nil {
		log.Fatal(err)
	}

	log.Println(result)
}
