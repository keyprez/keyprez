package router

import (
	"regexp"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

func getEndpointPath(path string) string {
	return ENDPOINT_PREFIX + path
}

type Route struct {
	path    string
	method  string
	handler LambdaHandler
}

func (r *Route) GetPath() string {
	return r.path
}

func (r *Route) GetMethod() string {
	return r.method
}

func (r *Route) SetHandler(handler LambdaHandler) *Route {
	r.handler = handler
	return r
}

func (r *Route) getMatcher() string {
	m := regexp.MustCompile(`\{[a-z]+\}`)
	matcher := m.ReplaceAllString(r.path, "([a-z0-9-]+)")
	return "^" + matcher + "$"
}

func (r *Route) IsMatch(partialUrl []byte) bool {
	checkMatch := regexp.MustCompile(r.getMatcher())
	if checkMatch.Match(partialUrl) {
		return true
	}
	return false
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

func NewRoute(path, method string, handler LambdaHandler) *Route {
	return &Route{path: getEndpointPath(path), method: method, handler: handler}
}
