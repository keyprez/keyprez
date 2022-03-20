package router

import "github.com/aws/aws-lambda-go/events"

func blankApiResponse(status int) *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		StatusCode:      status,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            "",
		IsBase64Encoded: false,
	}
}

func ReturnBlank200() (*events.APIGatewayProxyResponse, error) {
	return blankApiResponse(200), nil
}

func statusWithBody(status int, data string) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode:      status,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            data,
		IsBase64Encoded: false,
	}, nil
}

func Return200(data string) (*events.APIGatewayProxyResponse, error) {
	return statusWithBody(200, string(data))
}

func Return201(data string) (*events.APIGatewayProxyResponse, error) {
	return statusWithBody(201, string(data))
}

func Return204() (*events.APIGatewayProxyResponse, error) {
	return blankApiResponse(204), nil
}

func Return400() (*events.APIGatewayProxyResponse, error) {
	return blankApiResponse(400), nil
}

func Return404() (*events.APIGatewayProxyResponse, error) {
	return blankApiResponse(404), nil
}

func Return500() (*events.APIGatewayProxyResponse, error) {
	return blankApiResponse(500), nil
}
