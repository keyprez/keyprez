package router

import (
	"fmt"
	"regexp"
	"strings"

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

func (r *Route) GetMatcher() string {
	m := regexp.MustCompile(`\{[a-z]+\}`)
	matcher := m.ReplaceAllString(r.path, "([a-z0-9-]+)")
	return "^" + matcher + "$"
}

func (r *Route) SetUrlParameters(url string, request *events.APIGatewayProxyRequest) {
	m := regexp.MustCompile(`\{[a-z]+\}`)
	parameters := make(map[string]string)
	urlChunks := strings.Split(url, "/")
	pathChunks := strings.Split(r.path, "/")

	for index, value := range pathChunks {
		if !m.Match([]byte(value)) {
			continue
		}

		key := value[1 : len(value)-1]
		parameters[key] = urlChunks[index]
	}

	request.PathParameters = parameters
}

type Router struct {
	getRoutes  []*Route
	postRoutes []*Route
}

func (r *Router) Get(path string, handler lambdaHandler) *Route {
	route := &Route{path: path}
	r.getRoutes = append(r.getRoutes, route)
	return route.SetHandler(handler)
}

func (r *Router) Post(path string, handler lambdaHandler) *Route {
	route := &Route{path: path}
	r.postRoutes = append(r.postRoutes, route)
	return route.SetHandler(handler)
}

func (r *Router) GetHandler() lambdaHandler {
	return func(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
		var potentialRoutes []*Route
		fmt.Println(request.HTTPMethod)
		switch request.HTTPMethod {
		case "GET":
			potentialRoutes = r.getRoutes
			break
		case "POST":
			potentialRoutes = r.postRoutes
			break
		case "OPTIONS":
			return &events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers: map[string]string{
					"Content-Type":                 "application/json",
					"Access-Control-Allow-Methods": "HEAD,GET,PUT,DELETE,OPTIONS",
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
				checkMatch := regexp.MustCompile(route.GetMatcher())
				if checkMatch.Match(uriWithoutQuery) {
					route.SetUrlParameters(request.Path, &request)
					return route.handler(request)
				}
			}
		}

		return &events.APIGatewayProxyResponse{
			StatusCode:      404,
			Headers:         map[string]string{"Content-Type": "application/json"},
			Body:            "",
			IsBase64Encoded: false,
		}, nil
	}
}
