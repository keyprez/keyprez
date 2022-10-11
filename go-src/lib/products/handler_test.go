package products

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductsResponse struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       int32              `json:"price"`
	PriceID     string             `json:"priceID"`
	Stock       int32              `json:"stock"`
}

func TestProductsHandler(t *testing.T) {
	request := events.APIGatewayProxyRequest{
		Body: "{}",
	}
	msg, _ := ProductsHandler(request)

	var response []ProductsResponse
	json.Unmarshal([]byte(msg.Body), &response)

	if len(response) == 0 {
		t.Fatal("Comments should not be empty")
	}
}
