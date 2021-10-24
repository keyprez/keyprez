package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Comment struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Date    time.Time          `bson:"date,omitempty"`
	Name    string             `bson:"name,omitempty"`
	Email   string             `bson:"email,omitempty"`
	Text    string             `bson:"text,omitempty"`
	MovieId primitive.ObjectID `bson:"movie_id,omitempty"`
}

func productsHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	godotenv.Load("../../.env")

	atlasUri := os.Getenv("ATLAS_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(atlasUri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	db := os.Getenv("ATLAS_DB")
	commentsCollection := client.Database(db).Collection("comments")

	var comments []Comment

	selectOpts := options.Find()
	selectOpts.SetLimit(3)

	cursor, err := commentsCollection.Find(ctx, bson.M{"email": "john_bishop@fakegmail.com"}, selectOpts)
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &comments); err != nil {
		panic(err)
	}
	data, _ := json.Marshal(comments)
	return &events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           map[string]string{"Content-Type": "application/json"},
		MultiValueHeaders: http.Header{"Set-Cookie": {"Ding", "Ping"}},
		Body:              string(data),
		IsBase64Encoded:   false,
	}, nil
}
