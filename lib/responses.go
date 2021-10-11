package lib

import (
	"github.com/aws/aws-lambda-go/events"
)

//OkResponse is used to be CORS friendly response to the front-end
func OkResponse(body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: body, StatusCode: 200, Headers: corsHeaders()}
}

//EmptyOkResponse is used to be CORS friendly response to the front-end
func EmptyOkResponse() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: "", StatusCode: 200, Headers: corsHeaders()}
}

//CreatedResponse is setup for custom response that is CORS friendly.
func CreatedResponse() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: "", StatusCode: 201, Headers: corsHeaders()}
}

//NotFoundResponse is setup for responses of type not found.
func NotFoundResponse() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: "", StatusCode: 404, Headers: corsHeaders()}
}

func corsHeaders() map[string]string {
	h := make(map[string]string)
	h["Access-Control-Allow-Headers"] = "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token"
	h["Access-Control-Allow-Methods"] = "GET, POST, PUT, DELETE, OPTIONS"
	h["Access-Control-Allow-Origin"] = "*"

	return h
}
