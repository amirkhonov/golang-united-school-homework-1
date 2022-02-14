package main

import (
	"testing"
)

func TestHelloEmoji(t *testing.T) {
	if HelloEmoji("Hello") == "Hello 🗺️ " {
		t.Errorf("Sum was incorrect, got")
	}
}
