package customers

import (
	"encoding/json"
	"log"

	"github.com/keyprez/keyprez/go-src/lib/repository"
	"github.com/keyprez/keyprez/go-src/lib/repository/models"
	"github.com/keyprez/keyprez/go-src/lib/router"
	"github.com/keyprez/keyprez/go-src/lib/utils"

	"github.com/aws/aws-lambda-go/events"
)

func CreateCustomerHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var incoming *models.Customer
	err := json.Unmarshal([]byte(request.Body), &incoming)
	if err != nil {
		log.Printf("Failed parsing request body: %s", err)
		return router.Return400()
	}

	if !incoming.IsValid() {
		log.Print("Invalid customer")
		return router.Return400()
	}

	existing, err := repository.GetCustomerByEmail(incoming.Email)
	if err != nil {
		log.Printf("Could not fetch customer: %s", err)
		return router.Return400()
	}

	if existing != nil && existing.HasValidID() {
		if repository.MatchExisting(existing, incoming) {
			data, _ := json.Marshal(existing.StripeID)
			return router.Return200(string(data))
		}

		// Update customer in MongoDB
		err := repository.UpdateCustomer(existing.ID, incoming)
		if err != nil {
			log.Printf("Failed updating customer(MongoDB) with email %s", incoming.Email)
			return router.Return500()
		}

		data, _ := json.Marshal(existing.StripeID)

		// Update customer in Stripe
		_, stripeErr := utils.UpdateCustomer(existing.StripeID, incoming)
		if stripeErr != nil {
			log.Printf("Failed updating customer(Stripe) with email %s", incoming.Email)
			return router.Return500()
		}

		return router.Return200(string(data))
	}

	// Create customer in Stripe
	c, err := utils.CreateCustomer(incoming)
	if err != nil {
		return router.Return500()
	}

	response := &createResponse{
		ID: c.ID,
	}

	// Create customer in Mongo
	cust, err := repository.CreateCustomer(incoming, response.ID)
	if err != nil {
		return router.Return400()
	}

	if cust.HasValidID() {
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
