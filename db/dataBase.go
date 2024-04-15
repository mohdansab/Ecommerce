package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	User     *mongo.Collection
	Product  *mongo.Collection
	Cart     *mongo.Collection
	Mongo    *mongo.Client
	Database *mongo.Database
)

func ConnectDB() {
	Mongo_URL := "mongodb://127.0.0.1:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(Mongo_URL))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	Mongo = client
	Database = client.Database("crudSoib") // Initialize Database variable

	User = Database.Collection("user")
	Cart = Database.Collection("cart")
	Product = Database.Collection("product")

	fmt.Println("Connected to MongoDB")
}
