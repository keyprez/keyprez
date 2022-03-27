package orders

import (
	"encoding/json"
	"log"

	"github.com/keyprez/keyprez/go-src/lib/repository"
	"github.com/keyprez/keyprez/go-src/lib/utils"
	"github.com/stripe/stripe-go/v72"

	"github.com/keyprez/keyprez/go-src/lib/router"

	"github.com/aws/aws-lambda-go/events"
)

func CreateCheckoutSessionHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var requestBody createCheckoutSessionRequest
	if err := json.Unmarshal([]byte(request.Body), &requestBody); err != nil {
		return router.Return400()
	}

	s, err := utils.CreateCheckoutSession(requestBody.PriceID, requestBody.CustomerStripeID)
	if err != nil {
		return router.Return500()
	}

	response := &createCheckoutSessionResponse{
		ID: s.ID,
	}
	data, err := json.Marshal(response)
	if err != nil {
		return router.Return500()
	}

	return router.Return200(string(data))
}

func RetrieveSessionHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var requestBody retrieveSessionRequest
	if err := json.Unmarshal([]byte(request.Body), &requestBody); err != nil {
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

func WebhookHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	key := request.QueryStringParameters["key"]
	publishableKey := utils.GetEnvVar("STRIPE_PUBLISHABLE_KEY")
	if key != publishableKey {
		log.Println("Invalid or missing key")
		return router.Return400()
	}

	header, present := request.Headers["Stripe-Signature"]
	if !present {
		log.Println("Missing Stripe-Signature header")
		return router.Return400()

	}

	timestamp, err := utils.VerifyWebhookHeader(header)
	if err != nil {
		log.Println(err)
		return router.Return400()
	}

	event := stripe.Event{}
	if err := json.Unmarshal([]byte(request.Body), &event); err != nil {
		log.Printf("Error while parsing event. %v\n", err.Error())
		return router.Return500()
	}

	if event.Type != "checkout.session.completed" {
		log.Println("Only checkout.session.completed event is supported")
		return router.Return400()
	}

	if err := utils.VerifyTimestamp(event.Created, timestamp); err != nil {
		log.Printf("Error verifying webhook timestamp: %s\n", err)
		return router.Return400()
	}

	var checkoutSession stripe.CheckoutSession
	if err := json.Unmarshal(event.Data.Raw, &checkoutSession); err != nil {
		log.Printf("Error parsing webhook JSON: %v\n", err)
		return router.Return400()
	}

	s, sErr := utils.RetrieveSession(checkoutSession.ID)
	if sErr != nil {
		log.Printf("Error retrieving checkout session with ID %s: %v\n", checkoutSession.ID, sErr)
		return router.Return500()
	}

	sessionData := s.LineItems.Data
	if len(sessionData) > 1 {
		log.Println("Max one product is allowed")
		return router.Return500()
	}

	product := sessionData[0]

	updatedProduct, updatedProductErr := repository.UpdateProductStock(product.Price.ID, product.Quantity)
	if updatedProductErr != nil {
		log.Printf("Error updating stock amout for product %s: %v\n", product.Price.ID, updatedProductErr)
		return router.Return500()
	}

	data, dataErr := json.Marshal(updatedProduct)
	if dataErr != nil {
		return router.Return500()
	}

	// TODO: send confirmation email to the customer

	return router.Return200(string(data))
}

func SetupRouter() router.Router {
	r := router.Router{}
	r.Post("/orders/checkout", CreateCheckoutSessionHandler)
	r.Post("/orders/success", RetrieveSessionHandler)
	r.Post("/orders/webhook", WebhookHandler)

	return r
}
