package main

import (
	"testing"
)

func TestQuote(t *testing.T) {
	want := "Hello, world."
	if got := getQuote(); got != want {
		t.Errorf("TestHello() = %q, want %q", got, want)
	}
}

func TestNewerQuote(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := getNewerQuote(); got != want {
		t.Errorf("getNewerQuote() = %q, want %q", got, want)
	}
}
