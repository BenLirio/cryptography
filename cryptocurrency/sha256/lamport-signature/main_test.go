package main

import (
	"testing"
	"fmt"
)

func TestSign(t *testing.T) {
	signature, err := sign("key", "sample message")
	if err != nil {
		t.Errorf(fmt.Sprintf("%s", err))
	}
	if signature == nil {
		t.Errorf("NIL signature")
	}
	if len(signature) != 256 * 32 {
		t.Errorf("Expected signature length to be %d, got %d", 256 * 32, len(signature))
	}
}
