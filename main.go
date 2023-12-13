package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

type Response struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

func HandleRequest(ctx context.Context, event *MyEvent) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}

	if event.Name == "" {
		event.Name = "world"
	}

	response := Response{
		StatusCode: 200,
		Body:       fmt.Sprintf("Hello, %s!", event.Name),
	}

	bytes, _ := json.Marshal(response)
	message := string(bytes)
	return &message, nil
}

func main() {
	lambda.Start(HandleRequest)
}
