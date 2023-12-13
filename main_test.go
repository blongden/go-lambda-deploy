package main

import (
	"context"
	"testing"
)

func TestHandler(t *testing.T) {
	event := MyEvent{Name: "world"}
	ctx := context.TODO()

	message, _ := HandleRequest(ctx, &event)

	if *message != "Hello, world!" {
		t.Errorf("Unexpected output: %q", *message)
	}
}
