package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	id       int
	user_id  int
	name     string
	age      string
	avatar   string
	url      string
	gender   string
	detail   string
	follwers string
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://192.168.1.3:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("go-mongo").Collection("user")

	james := User{51, 33321, "James", "32", "a1", "u1", "n", "d1", ""}
	wade := User{52, 3322, "wade", "35", "a2", "u2", "n", "d2", ""}

	trainers := []interface{}{james, wade}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

}
