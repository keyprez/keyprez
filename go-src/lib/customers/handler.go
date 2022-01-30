package customers

import (
	"encoding/json"
	"fmt"

	"github.com/keyprez/keyprez/go-src/lib/router"
	"github.com/keyprez/keyprez/go-src/lib/utils"

	"github.com/keyprez/keyprez/go-src/lib/models"

	"github.com/aws/aws-lambda-go/events"
)

type createRequest struct {
	Email string `json:"email"`
}

type createResponse struct {
	ID string `json:"id"`
}

func CreateCustomerHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var requestBody createRequest
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		return router.Return400()
	}

	if !utils.IsValidEmail(requestBody.Email) {
		return router.Return400()
	}

	customer, err := models.GetCustomerByEmail(requestBody.Email)
	if err != nil {
		return router.Return400()
	}

	if customer.IsValid() {
		data, _ := json.Marshal(customer)
		return &events.APIGatewayProxyResponse{
			StatusCode:      200,
			Headers:         map[string]string{"Content-Type": "application/json"},
			Body:            string(data),
			IsBase64Encoded: false,
		}, nil
	}

	c, err := utils.CreateCustomer(requestBody.Email)
	if err != nil {
		return router.Return500()
	}

	response := &createResponse{
		ID: c.ID,
	}

	fmt.Println("Utils customer created")
	cust, err := models.CreateCustomer(requestBody.Email, response.ID)
	if err != nil {
		return router.Return400()
	}
	data, err := json.Marshal(cust)
	fmt.Println("Models customer created")

	return &events.APIGatewayProxyResponse{
		StatusCode:      201,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            string(data),
		IsBase64Encoded: false,
	}, nil
}

func SetupRouter() router.Router {
	r := router.Router{}
	r.Post("/customers", CreateCustomerHandler)
	return r
}
