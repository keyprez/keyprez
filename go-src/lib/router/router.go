package router

import (
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type Router struct {
	getRoutes    []*Route
	postRoutes   []*Route
	putRoutes    []*Route
	deleteRoutes []*Route
}

func (r *Router) Get(path string, handler LambdaHandler) *Route {
	route := NewRoute(path, http.MethodGet, handler)
	r.getRoutes = append(r.getRoutes, route)
	return route
}

func (r *Router) Post(path string, handler LambdaHandler) *Route {
	route := NewRoute(path, http.MethodPost, handler)
	r.getRoutes = append(r.getRoutes, route)
	return route
}

func (r *Router) Put(path string, handler LambdaHandler) *Route {
	route := NewRoute(path, http.MethodPut, handler)
	r.getRoutes = append(r.getRoutes, route)
	return route
}

func (r *Router) Delete(path string, handler LambdaHandler) *Route {
	route := NewRoute(path, http.MethodDelete, handler)
	r.getRoutes = append(r.getRoutes, route)
	return route
}

func (r *Router) GetAllRoutes() []*Route {
	combinedRoutes := []*Route{}
	combinedRoutes = append(r.getRoutes, r.postRoutes...)
	combinedRoutes = append(combinedRoutes, r.putRoutes...)
	combinedRoutes = append(combinedRoutes, r.deleteRoutes...)
	return combinedRoutes
}

func (r *Router) GetHandler() LambdaHandler {
	return func(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
		var potentialRoutes []*Route
		log.Println("Received request", request.HTTPMethod, request.Path)

		switch request.HTTPMethod {
		case "GET":
			potentialRoutes = r.getRoutes
			break
		case "POST":
			potentialRoutes = r.postRoutes
			break
		case "PUT":
			potentialRoutes = r.putRoutes
			break
		case "DELETE":
			potentialRoutes = r.deleteRoutes
			break
		case "OPTIONS":
			return &events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers: map[string]string{
					"Content-Type":                 "application/json",
					"Access-Control-Allow-Methods": "HEAD,GET,POST,PUT,DELETE,OPTIONS",
					"Access-Control-Allow-Origin":  "*",
					"Access-Control-Allow-Headers": "*",
				},
				Body:            "",
				IsBase64Encoded: false,
			}, nil
		}

		if len(potentialRoutes) >= 0 {
			uriWithoutQuery := []byte(strings.Split(request.Path, "?")[0])
			for _, route := range potentialRoutes {
				if route.IsMatch(uriWithoutQuery) {
					route.SetUrlParameters(request.Path, &request)
					return route.handler(request)
				}
			}
		}

		return Return404()
	}
}
