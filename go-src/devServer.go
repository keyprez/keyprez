package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"github.com/keyprez/keyprez/go-src/lib/products"
)

func productsHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Received request")
	response, err := products.ProductsHandler(events.APIGatewayProxyRequest{})
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(response.Body))
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/products", productsHandle)
	err := http.ListenAndServe(":8090", nil)

	if err != nil {
		fmt.Println(err)
	}
}
