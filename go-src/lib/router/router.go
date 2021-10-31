package router

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type lambdaHandler func(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)

type Route struct {
	path    string
	handler lambdaHandler
}

func (r *Route) SetHandler(handler lambdaHandler) *Route {
	r.handler = handler
	return r
}

type Router struct {
	routes []*Route
}

func (r *Router) Get(path string, handler lambdaHandler) *Route {
	route := &Route{}
	r.routes = append(r.routes, route)
	return route.SetHandler(handler)
}

func (r *Router) GetHandler() lambdaHandler {
	return func(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
		return &events.APIGatewayProxyResponse{
			StatusCode:        200,
			Headers:           map[string]string{"Content-Type": "application/json"},
			MultiValueHeaders: http.Header{"Set-Cookie": {"Ding", "Ping"}},
			Body:              "Hello",
			IsBase64Encoded:   false,
		}, nil
	}
}
