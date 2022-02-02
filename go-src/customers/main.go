package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/keyprez/keyprez/go-src/lib/customers"
)

func main() {
	customersRouter := customers.SetupRouter()
	lambda.Start(customersRouter.GetHandler())
}
