package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/aws/aws-lambda-go/events"

	"github.com/keyprez/keyprez/go-src/lib/orders"
	"github.com/keyprez/keyprez/go-src/lib/products"
	"github.com/keyprez/keyprez/go-src/lib/router"
)

type ProductRequestPost struct {
	A string `json:"a"`
}

func transformRequest(req *http.Request) events.APIGatewayProxyRequest {
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
		Path:                            req.RequestURI,
	}
}

func setupAndRespond(w http.ResponseWriter, req *http.Request, setupRouter func() router.Router, log string) {
	fmt.Println(log)
	router := setupRouter()
	response, err := router.GetHandler()(transformRequest(req))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "HEAD,GET,PUT,DELETE,OPTIONS")
	w.Write([]byte(response.Body))
}

func productsHandle(w http.ResponseWriter, req *http.Request) {
	setupAndRespond(w, req, products.SetupRouter, "Received product request")
}

func ordersHandle(w http.ResponseWriter, req *http.Request) {
	setupAndRespond(w, req, orders.SetupRouter, "Received order request")
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
	// Todo: Use the routers to create the handlers for mux
	r.HandleFunc("/.netlify/functions/products", productsHandle)
	r.HandleFunc("/.netlify/functions/products/{key}", productsHandle)
	r.HandleFunc("/.netlify/functions/orders/checkout", ordersHandle)

	fmt.Println("Listening on port 8090")
	err := http.ListenAndServe(":8090", r)

	if err != nil {
		panic(err)
	}
}
