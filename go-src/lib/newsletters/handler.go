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
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
	}

	_, emailErr := mail.ParseAddress(requestBody.Email)
	if emailErr != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode:      400,
			Headers:         map[string]string{"Content-Type": "application/json"},
			Body:            "Invalid email address",
			IsBase64Encoded: false,
		}, nil
	}

	if requestBody.Email == "" {
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
	}

	subscription, err := models.GetNewsletterSubscriptionByEmail(requestBody.Email)
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
	}

	if subscription != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
	}

	success, err := models.CreateNewsletterSubscription(requestBody.Email)
	if !success || err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func SetupRouter() router.Router {
	r := router.Router{}
	r.Post("/.netlify/functions/newsletters", CreateSubscriptionHandler)

	return r
}
