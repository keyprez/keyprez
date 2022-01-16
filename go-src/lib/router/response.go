package router

import "github.com/aws/aws-lambda-go/events"

func blankApiResponse(status int) *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}
}

func ReturnBlank200() (*events.APIGatewayProxyResponse, error) {
	return blankApiResponse(200), nil
}

func Return200(data string) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            data,
		IsBase64Encoded: false,
	}, nil
}

func Return201() (*events.APIGatewayProxyResponse, error) {
	return blankApiResponse(201), nil
}

func Return400() (*events.APIGatewayProxyResponse, error) {
	return blankApiResponse(400), nil
}

func Return500() (*events.APIGatewayProxyResponse, error) {
	return blankApiResponse(500), nil
}
