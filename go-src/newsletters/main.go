package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/keyprez/keyprez/go-src/lib/newsletters"
)

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	newslettersRouter := newsletters.SetupRouter()
	lambda.Start(newslettersRouter.GetHandler())
}
