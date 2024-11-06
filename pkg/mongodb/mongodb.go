package mongodb

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

func Connect() {
	// getting mongodb uri
	var uri = os.Getenv("MONGODB_URI")
	// Connection to database
	var ctx = context.Background()
	var options = options.Client().ApplyURI(uri)
	var client, err = mongo.Connect(ctx, options)

	if err != nil {
		fmt.Println("MONGODB CONNECTION FAILED")
	}

	// Database
	Db = client.Database("voting-tool")
}
