package products

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductsResponse struct {
	ID          primitive.ObjectID `json:"ID"`
	Name        string             `json:"Name"`
	Description string             `json:"Description"`
	Price       int32              `json:"Price"`
	PriceID     string             `json:"PriceID"`
	Stock       int32              `json:"Stock"`
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
