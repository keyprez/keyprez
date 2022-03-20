package newsletters

import (
	"encoding/json"

	"github.com/keyprez/keyprez/go-src/lib/repository"
	"github.com/keyprez/keyprez/go-src/lib/router"
	"github.com/keyprez/keyprez/go-src/lib/utils"

	"github.com/aws/aws-lambda-go/events"
)

func CreateSubscriptionHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var requestBody createRequest
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		return router.Return400()
	}

	if !utils.IsValidEmail(requestBody.Email) {
		return router.Return400()
	}

	subscription, err := repository.GetNewsletterSubscriptionByEmail(requestBody.Email)
	if err != nil {
		return router.Return400()
	}

	if subscription.IsValid() {
		return router.ReturnBlank200()
	}

	success, err := repository.CreateNewsletterSubscription(requestBody.Email)
	if !success || err != nil {
		return router.Return400()
	}

	return router.Return201("Subscription created")
}

func UnsubscribeSubscriptionHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	email := request.QueryStringParameters["email"]
	if !utils.IsValidEmail(email) {
		return router.Return400()
	}

	subscription, err := repository.GetNewsletterSubscriptionByEmail(email)
	if err != nil {
		return router.Return400()
	}

	if ok, err := repository.UnsubscribeNewsletterSubscription(subscription); err != nil || !ok {
		return router.Return500()
	}

	return router.Return204()
}

func SetupRouter() router.Router {
	r := router.Router{}
	r.Post("/newsletters", CreateSubscriptionHandler)
	r.Get("/newsletters/unsubscribe", UnsubscribeSubscriptionHandler)
	return r
}
