package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DBSet()

func DBSet() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("failed to connect to mongoDB")
		return nil
	}

	fmt.Println("connected to mongoDB")
	return client
}

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("goecommerce").Collection(collectionName)
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("goecommerce").Collection(collectionName)
}
