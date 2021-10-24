package products

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentResponse struct {
	ID      primitive.ObjectID `json:"ID"`
	Date    time.Time          `json:"Date"`
	Name    string             `json:"Name"`
	Email   string             `json:"Email"`
	Text    string             `json:"Text"`
	MovieId primitive.ObjectID `json:"MovieId"`
}

func TestProductsHandler(t *testing.T) {
	request := events.APIGatewayProxyRequest{
		Body: "{}",
	}
	msg, _ := ProductsHandler(request)

	var comments []CommentResponse
	json.Unmarshal([]byte(msg.Body), &comments)

	if len(comments) == 0 {
		t.Fatal("Comments should not be empty")
	}
}
