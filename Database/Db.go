package Database

import (
	"context"
	"time"

	"github.com/amadou-toure/groove_api/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var Mg models.MongoInstance
const dbName="groove_DB"
const mongoURI="mongodb://localhost:27017/" + dbName

func Connect() error {
    // Create a new MongoDB client
    client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
    if err != nil {
        return err
    }

    // Create a context with a timeout
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel() // Ensure the context is canceled to release resources

    // Connect to the MongoDB server
    err = client.Connect(ctx)
    if err != nil {
        return err
    }

    // Verify the connection
    err = client.Ping(ctx, nil)
    if err != nil {
        return err
    }

    // Access the database
    db := client.Database(dbName)

    // Assign the MongoDB client and database to the global variable
    Mg = models.MongoInstance{Client: client, Db: db}

    return nil
}

