package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello Go"

	val := hello()
	if val != expected {
		t.Errorf("got %q, want %q", val, expected)
	}
}
