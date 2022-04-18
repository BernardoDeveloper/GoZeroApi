package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	fmt.Println("Connected to MongoDB!")
	return nil
}

func StartConnect() {
	// Connect to MongoDB
	client, ctx, cancel, err := connect("mongodb+srv://bernardoDeveloper:Bernardo300@cluster0.zcewb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	if err != nil {
		log.Fatal(err)
	}

	defer close(client, ctx, cancel)
	ping(client, ctx)
}

func GetMongoClient() (*mongo.Client, context.Context) {
	client, ctx, _, _ := connect("mongodb+srv://bernardoDeveloper:Bernardo300@cluster0.zcewb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	return client, ctx
}
