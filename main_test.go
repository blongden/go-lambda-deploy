package main

import (
	"context"
	"encoding/json"
	"testing"
)

func TestHandler(t *testing.T) {
	event := MyEvent{Name: "golang"}
	message, _ := HandleRequest(context.Background(), &event)
	response := Response{}

	json.Unmarshal([]byte(*message), &response)

	if response.Body != "Hello, golang!" {
		t.Errorf("Unexpected output: %q", *message)
	}
}

func TestHandlerDefault(t *testing.T) {
	event := MyEvent{}
	message, _ := HandleRequest(context.Background(), &event)
	response := Response{}

	json.Unmarshal([]byte(*message), &response)

	if response.Body != "Hello, world!" {
		t.Errorf("Unexpected output: %q", *message)
	}
}
