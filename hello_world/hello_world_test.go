package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Eddie")
	want := "Hello, Eddie"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}