package newsletters

import (
	"encoding/json"
	"net/mail"

	"github.com/keyprez/keyprez/go-src/lib/router"

	"github.com/keyprez/keyprez/go-src/lib/models"

	"github.com/aws/aws-lambda-go/events"
)

type createRequest struct {
	Email string `json:"email"`
}

func CreateSubscriptionHandler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var requestBody createRequest
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		return router.Return400()
	}

	_, emailErr := mail.ParseAddress(requestBody.Email)
	if emailErr != nil {
		return router.Return400()
	}

	if requestBody.Email == "" {
		return router.Return400()
	}

	subscription, err := models.GetNewsletterSubscriptionByEmail(requestBody.Email)
	if err != nil {
		return router.Return400()
	}

	if subscription != nil {
		return router.ReturnBlank200()
	}

	success, err := models.CreateNewsletterSubscription(requestBody.Email)
	if !success || err != nil {
		return router.Return400()
	}

	return router.Return201()
}

func SetupRouter() router.Router {
	r := router.Router{}
	r.Post("/.netlify/functions/newsletters", CreateSubscriptionHandler)

	return r
}
