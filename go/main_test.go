package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	res := EvenOrOdd(10)
	if res != "Even" {
		t.Errorf("Expected Even, got %s", res)
	}
}
