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

func parseMyEventJson(req *MyEvent, body string) error {
	if body != "" {
		err := json.Unmarshal([]byte(body), &req)

		if err != nil {
			return err
		}
	}
	return nil
}

func HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req := MyEvent{}

	parseMyEventJson(&req, event.Body)

	if req.Name == "" {
		req.Name = "world"
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("Goodbye, %s!", req.Name),
	}

	return response, nil
}

func main() {
	lambda.Start(HandleRequest)
}
