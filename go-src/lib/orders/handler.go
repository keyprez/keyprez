package orders

import (
	"encoding/json"

	"github.com/keyprez/keyprez/go-src/lib/utils"

	"github.com/keyprez/keyprez/go-src/lib/router"

	"github.com/aws/aws-lambda-go/events"
)

type createCheckoutSessionRequest struct {
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

	s, err := utils.CreateCheckoutSession(requestBody.PriceID, requestBody.CustomerStripeID)
	if err != nil {
		return router.Return500()
	}

	response := &createCheckoutSessionResponse{
		ID: s.ID,
	}
	data, _ := json.Marshal(response)

	return router.Return200(string(data))
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
	if err != nil {
		return router.Return500()
	}

	return router.Return200(string(data))
}

func SetupRouter() router.Router {
	r := router.Router{}
	//r.Get("/.netlify/functions/orders", ProductsHandler)
	r.Post("/orders/checkout", CreateCheckoutSessionHandler)
	r.Post("/orders/success", RetrieveSessionHandler)

	return r
}
