package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	req, _ := json.Marshal(MyEvent{Name: "golang"})
	event := events.APIGatewayProxyRequest{Body: string(req)}
	message, _ := HandleRequest(context.Background(), &event)

	if message.Body != "Hello, golang!" {
		t.Errorf("Unexpected output: %q", message.Body)
	}
}

func TestHandlerDefault(t *testing.T) {
	req, _ := json.Marshal(MyEvent{})
	event := events.APIGatewayProxyRequest{Body: string(req)}
	message, _ := HandleRequest(context.Background(), &event)

	if message.Body != "Hello, world!" {
		t.Errorf("Unexpected output: %q", message.Body)
	}
}
