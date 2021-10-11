package lib

import (
	"github.com/aws/aws-lambda-go/events"
)

func Response200(body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: body, StatusCode: 200, Headers: corsHeaders()}
}

func Response404() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: "", StatusCode: 404, Headers: corsHeaders()}
}

func corsHeaders() map[string]string {
	h := make(map[string]string)
	h["Access-Control-Allow-Headers"] = "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token"
	h["Access-Control-Allow-Methods"] = "GET, POST, PUT, DELETE, OPTIONS"
	h["Access-Control-Allow-Origin"] = "*"

	return h
}
