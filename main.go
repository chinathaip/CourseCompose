package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chinathaip/coursecompose/router"
	"github.com/chinathaip/coursecompose/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	client, cancel, err := connectDB(os.Getenv("DB_CONNECTION"))
	if err != nil {
		log.Fatalln("Error connection to DB:", err)
	}
	mongoService := service.NewMongoService(client)

	h := router.NewHandler(mongoService)
	e := h.RegisterRoute()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	srv := http.Server{
		Addr:    ":1234",
		Handler: e,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	<-signals
	cancel()
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalln("Error shutting down: ", err)
	}
}

func connectDB(dbconn string) (*mongo.Client, context.CancelFunc, error) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(dbconn).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, cancel, err
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, cancel, err
	}

	return client, cancel, nil
}
