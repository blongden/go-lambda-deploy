package main

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func checkGreeting(t *testing.T, body string, name string) {
	if body != fmt.Sprintf("Hello, %s!", name) {
		t.Errorf("Unexpected output: %q", body)
	}
}

func TestHandler(t *testing.T) {
	req, _ := json.Marshal(MyEvent{Name: "golang"})
	event := events.APIGatewayProxyRequest{Body: string(req)}
	message, _ := HandleRequest(context.Background(), event)

	checkGreeting(t, message.Body, "golang")
}

func TestHandlerDefault(t *testing.T) {
	req, _ := json.Marshal(MyEvent{})
	event := events.APIGatewayProxyRequest{Body: string(req)}
	message, _ := HandleRequest(context.Background(), event)

	checkGreeting(t, message.Body, "world")
}
