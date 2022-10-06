package customers

import (
	"encoding/json"

	"github.com/keyprez/keyprez/go-src/lib/repository"
	"github.com/keyprez/keyprez/go-src/lib/router"
	"github.com/keyprez/keyprez/go-src/lib/utils"

	"github.com/aws/aws-lambda-go/events"
)

func CreateCustomerHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var requestBody utils.CreateCustomerRequest
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		return router.Return400()
	}

	if !utils.IsValidCustomer(requestBody) {
		return router.Return400()
	}

	customer, err := repository.GetCustomerByEmail(requestBody.Email)
	if err != nil {
		return router.Return400()
	}

	if customer != nil && customer.IsValid() {
		data, _ := json.Marshal(customer.StripeID)
		return router.Return200(string(data))
	}

	// Create customer in Stripe
	c, err := utils.CreateCustomer(requestBody)
	if err != nil {
		return router.Return500()
	}

	response := &createResponse{
		ID: c.ID,
	}

	// Create customer in Mongo
	cust, err := repository.CreateCustomer(requestBody, response.ID)
	if err != nil {
		return router.Return400()
	}

	if cust.IsValid() {
		data, _ := json.Marshal(cust.StripeID)
		return router.Return201(string(data))
	}

	return router.Return500()
}

func SetupRouter() router.Router {
	r := router.Router{}
	r.Post("/customers", CreateCustomerHandler)
	return r
}
