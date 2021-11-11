package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/keyprez/keyprez/go-src/lib/orders"
)

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	ordersRouter := orders.SetupRouter()
	lambda.Start(ordersRouter.GetHandler())
}
