package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := getQuote(); got != want {
		t.Errorf("TestHello() = %q, want %q", got, want)
	}
}

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := getNewerQuote(); got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
}
