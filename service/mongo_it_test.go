//go:build integration

package service

import (
	"context"
	"log"
	"testing"

	"github.com/chinathaip/coursecompose/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestGetAllCourses(t *testing.T) {
	var (
		client         = setup()
		service        = NewMongoService(client)
		expectedLength = 3
	)
	defer teardown(client)

	response := service.GetAllCourses(bson.M{})

	assert.Equal(t, expectedLength, len(response))
	assert.Equal(t, "ITE224", response[0].CourseCode)
	assert.Equal(t, "ITE102", response[1].CourseCode)
	assert.Equal(t, "ITE240", response[2].CourseCode)
}

// utility func
func setup() *mongo.Client {
	var (
		opts      = options.Client().ApplyURI("mongodb://db:27017")
		client, _ = mongo.NewClient(opts)
	)

	if err := client.Connect(context.TODO()); err != nil {
		log.Fatalln("Cannot connect to MongoDB", err)
	}

	seedDB(client)

	return client
}

func teardown(client *mongo.Client) {
	clearDB(client)
	client.Disconnect(context.Background())
}

func seedDB(client *mongo.Client) {
	col := client.Database(database).Collection(courses)
	courses := []model.Course{
		{
			Code: "ITE224",
			Name: "Intro to Data Science",
		},
		{
			Code: "ITE102",
			Name: "Discrete Math",
		},
		{
			Code: "ITE240",
			Name: "Operating System",
		},
	}

	docs := courseToBSON(courses)

	result, _ := col.InsertMany(context.TODO(), docs)
	log.Printf("seeded: %d items to mock db", len(result.InsertedIDs))
}

func clearDB(client *mongo.Client) {
	col := client.Database(database).Collection(courses)

	col.DeleteMany(context.Background(), bson.D{})
}

func courseToBSON(courses []model.Course) []interface{} {
	var docs []interface{}
	for _, course := range courses {
		doc, _ := bson.Marshal(course)
		docs = append(docs, doc)
	}

	return docs
}
