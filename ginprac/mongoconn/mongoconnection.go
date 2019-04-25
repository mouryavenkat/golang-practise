package mongoconn

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//Client ... A global connection Object for mongo client connection
var Client *mongo.Client

//ConnectToMongo ... Establish a new connection to mongo when the server starts up.
func ConnectToMongo() {
	Client, _ = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, err1 := context.WithTimeout(context.Background(), 10*time.Second)
	if err1 != nil {
	}
	// time.Second returns 1 second. This will wait for 10 seconds
	Client.Connect(ctx)
	err := Client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Unable to ping to mongo instance")
		// log.Fatalf will actually halt the current running process
	}
	log.Println("Mongo connection is successfully established at localhost:27017")
}
