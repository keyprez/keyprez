package orders

import (
	"encoding/json"
	"fmt"

	"github.com/keyprez/keyprez/go-src/lib/models"
	"github.com/keyprez/keyprez/go-src/lib/utils"
	"github.com/stripe/stripe-go/v72"

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

func WebhookHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod != "POST" {
		fmt.Println("Only POST method is allowed")
		return router.Return400()
	}

	signature, present := request.Headers["Stripe-Signature"]
	if !present {
		fmt.Printf("Missing Stripe-Signature header\n")
		return router.Return400()
	}

	if err := utils.VerifySignature([]byte(request.Body), signature); err != nil {
		fmt.Printf("Error verifying webhook signature\n")
		return router.Return400()
	}

	event := stripe.Event{}
	if err := json.Unmarshal([]byte(request.Body), &event); err != nil {
		fmt.Printf("Error while parsing event. %v\n", err.Error())
		return router.Return500()
	}

	if event.Type != "checkout.session.completed" {
		return router.Return400()
	}

	var checkoutSession stripe.CheckoutSession
	if err := json.Unmarshal(event.Data.Raw, &checkoutSession); err != nil {
		fmt.Printf("Error parsing webhook JSON: %v\n", err)
		return router.Return400()
	}

	s, sErr := utils.RetrieveSession(checkoutSession.ID)
	if sErr != nil {
		fmt.Printf("Error retrieving checkout session with ID %s: %v\n", checkoutSession.ID, sErr)
		return router.Return500()
	}

	sessionData := s.LineItems.Data
	if len(sessionData) > 1 {
		fmt.Println("Max one product is allowed")
		return router.Return500()
	}

	product := sessionData[0]

	updatedProduct, updatedProductErr := models.UpdateProductStock(product.Price.ID, product.Quantity)
	if updatedProductErr != nil {
		fmt.Printf("Error updating stock amout for product %s: %v\n", product.Price.ID, updatedProductErr)
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
	//r.Get("/.netlify/functions/orders", ProductsHandler)
	r.Post("/orders/checkout", CreateCheckoutSessionHandler)
	r.Post("/orders/success", RetrieveSessionHandler)
	r.Post("/orders/webhook", WebhookHandler)

	return r
}
