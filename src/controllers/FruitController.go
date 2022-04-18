package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BernardoDeveloper/learngoapifromzero/config/database"
	"github.com/BernardoDeveloper/learngoapifromzero/src/models"
	"go.mongodb.org/mongo-driver/bson"
)

var client, _ = database.GetMongoClient()
var collection = client.Database("learngoapi").Collection("fruits")
var f models.Fruit

func ShowFruits(w http.ResponseWriter, r *http.Request) {
	filter := bson.D{{}}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
		w.Write(output)
	}

}

func CreateFruit(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&f)

	value, err := collection.InsertOne(context.TODO(), f)
	result, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	w.Write(result)
}
