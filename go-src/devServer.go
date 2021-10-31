package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/aws/aws-lambda-go/events"

	"github.com/keyprez/keyprez/go-src/lib/products"
)

type ProductRequestPost struct {
	A string `json:"a"`
}

func transformRequest(req *http.Request) events.APIGatewayProxyRequest {
	fmt.Println(req.RequestURI)
	pathParams := mux.Vars(req)
	defer req.Body.Close()

	var mapInterface ProductRequestPost
	_ = json.NewDecoder(req.Body).Decode(&mapInterface)

	queryStringParams := make(map[string]string)
	multiValueQueryStringParams := make(map[string][]string)
	for key, val := range req.URL.Query() {
		if len(val) > 1 {
			multiValueQueryStringParams[key] = val
		} else {
			queryStringParams[key] = val[0]
		}
	}

	return events.APIGatewayProxyRequest{
		HTTPMethod:                      req.Method,
		Headers:                         headers(req),
		PathParameters:                  pathParams,
		QueryStringParameters:           queryStringParams,
		MultiValueQueryStringParameters: multiValueQueryStringParams,
	}
}

func productsHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Received request")
	response, err := products.ProductsHandler(transformRequest(req))
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(response.Body))
}

func headers(req *http.Request) map[string]string {
	m := make(map[string]string)
	for name, headers := range req.Header {
		for _, h := range headers {
			m[name] = h
		}
	}
	return m
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/.netlify/functions/products", productsHandle)
	r.HandleFunc("/.netlify/functions/products/{key}", productsHandle)

	fmt.Println("Listening on port 8090")
	err := http.ListenAndServe(":8090", r)

	if err != nil {
		panic(err)
	}
}
