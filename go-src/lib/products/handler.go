package products

import (
	"encoding/json"

	"github.com/keyprez/keyprez/go-src/lib/router"

	"github.com/keyprez/keyprez/go-src/lib/models"

	"github.com/aws/aws-lambda-go/events"
)

func ProductHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	products, err := models.GetProduct()
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(products)

	return router.Return200(string(data))
}

func ProductsHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	products, err := models.GetProducts()
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(products)

	return router.Return200(string(data))
}

func SetupRouter() router.Router {
	r := router.Router{}
	r.Get("/products", ProductsHandler)
	r.Get("/products/{name}", ProductHandler)

	return r
}
