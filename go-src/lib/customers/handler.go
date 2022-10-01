package customers

import (
	"encoding/json"

	"github.com/keyprez/keyprez/go-src/lib/repository"
	"github.com/keyprez/keyprez/go-src/lib/router"
	"github.com/keyprez/keyprez/go-src/lib/utils"

	"github.com/aws/aws-lambda-go/events"
)

func CreateCustomerHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var requestBody createRequest
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		return router.Return400()
	}

	if !utils.IsValidEmail(requestBody.Email) {
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

	c, err := utils.CreateCustomer(requestBody.Email)
	if err != nil {
		return router.Return500()
	}

	response := &createResponse{
		ID: c.ID,
	}

	cust, err := repository.CreateCustomer(requestBody.Email, response.ID)
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
