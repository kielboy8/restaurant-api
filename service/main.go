package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/kielboy8/restaurant-api/lib"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return lib.NotFoundResponse(), nil
}

func main() {
	lambda.Start(handler)
}
