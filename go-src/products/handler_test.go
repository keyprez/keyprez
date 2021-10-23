package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestProductsHandler(t *testing.T) {
	request := events.APIGatewayProxyRequest{
		Body: "{}",
	}
	msg, _ := handler(request)

	expected := "Hello, World!"
	if msg.Body != expected {
		t.Fatal("Expected " + expected + ", got " + msg.Body)
	}
}
