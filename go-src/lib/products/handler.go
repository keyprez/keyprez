package products

import (
	"encoding/json"
	"net/http"

	"github.com/keyprez/keyprez/go-src/lib/router"

	"github.com/keyprez/keyprez/go-src/lib/models"

	"github.com/aws/aws-lambda-go/events"
)

type getResponse struct {
	Data []models.Product
	Path string
}

func ProductHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	products, err := models.GetProducts()
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(getResponse{
		Data: products,
		Path: request.Path,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           map[string]string{"Content-Type": "application/json"},
		MultiValueHeaders: http.Header{"Set-Cookie": {"Ding", "Ping"}},
		Body:              string(data),
		IsBase64Encoded:   false,
	}, nil
}

func ProductsHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	products, err := models.GetProducts()
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(getResponse{
		Data: products,
		Path: request.Path,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           map[string]string{"Content-Type": "application/json"},
		MultiValueHeaders: http.Header{"Set-Cookie": {"Ding", "Ping"}},
		Body:              string(data),
		IsBase64Encoded:   false,
	}, nil
}

func SetupRouter() router.Router {
	r := router.Router{}
	r.Get("/.netlify/functions/products", ProductsHandler)
	r.Get("/.netlify/functions/products/{id}", ProductHandler)

	return r
}
