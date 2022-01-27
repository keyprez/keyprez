package orders

import (
	"encoding/json"

	"github.com/keyprez/keyprez/go-src/lib/utils"

	"github.com/keyprez/keyprez/go-src/lib/router"

	"github.com/aws/aws-lambda-go/events"
)

type createCheckoutSessionResponse struct {
	ID string `json:"id"`
}

func CreateCheckoutSessionHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	s, err := utils.CreateCheckoutSession()
	if err != nil {
		return router.Return500()
	}

	response := &createCheckoutSessionResponse{
		ID: s.ID,
	}
	data, err := json.Marshal(response)

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            string(data),
		IsBase64Encoded: false,
	}, nil
}

func SetupRouter() router.Router {
	r := router.Router{}
	//r.Get("/.netlify/functions/orders", ProductsHandler)
	r.Post("/orders/checkout", CreateCheckoutSessionHandler)

	return r
}
