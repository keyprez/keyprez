package products

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/keyprez/keyprez/go-src/lib/models"

	"github.com/aws/aws-lambda-go/events"
)

func ProductsHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	products, err := models.GetProducts()
	if err != nil {
		log.Fatal(err)
	}

	data, _ := json.Marshal(products)
	return &events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           map[string]string{"Content-Type": "application/json"},
		MultiValueHeaders: http.Header{"Set-Cookie": {"Ding", "Ping"}},
		Body:              string(data),
		IsBase64Encoded:   false,
	}, nil
}
