package conection

import (
    "context"
    "fmt"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConection(ctx context.Context) *mongo.Client {
    // Set up a context and options for the MongoDB client
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

    // Connect to the MongoDB server
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // Check the connection
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")

    return client
}
