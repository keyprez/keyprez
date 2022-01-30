package orders

import (
	"encoding/json"

	"github.com/keyprez/keyprez/go-src/lib/utils"

	"github.com/keyprez/keyprez/go-src/lib/router"

	"github.com/aws/aws-lambda-go/events"
)

type createCheckoutSessionRequest struct {
	ProductName      string `json:"productname"`
	PriceID          string `json:"priceid"`
	CustomerStripeID string `json:"customerstripeid"`
}

type createCheckoutSessionResponse struct {
	ID string `json:"id"`
}

type retrieveSessionRequest struct {
	SessionID string `json:"sessionid"`
}

func CreateCheckoutSessionHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var requestBody createCheckoutSessionRequest
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		return router.Return400()
	}

	s, err := utils.CreateCheckoutSession(requestBody.ProductName, requestBody.PriceID, requestBody.CustomerStripeID)
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

func RetrieveSessionHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var requestBody retrieveSessionRequest
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		return router.Return400()
	}

	s, err := utils.RetrieveSession(requestBody.SessionID)
	if err != nil {
		return router.Return500()
	}

	data, err := json.Marshal(s)

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
	r.Post("/orders/success", RetrieveSessionHandler)

	return r
}
