package router

import "github.com/aws/aws-lambda-go/events"

const ENDPOINT_PREFIX string = "/.netlify/functions"

type LambdaHandler func(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)
