package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

type Response struct {
	StatusCode      int    `json:"statusCode"`
	Body            string `json:"body"`
	IsBase64Encoded bool   `json:"isBase64Encoded"`
}

func HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	req := MyEvent{}

	if event.Body != "" {
		err := json.Unmarshal([]byte(event.Body), &req)

		if err != nil {
			return &events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "Unable to parse JSON body",
			}, nil
		}
	}

	if req.Name == "" {
		req.Name = "world"
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("Hello, %s!", req.Name),
	}

	return &response, nil
}

func main() {
	lambda.Start(HandleRequest)
}
