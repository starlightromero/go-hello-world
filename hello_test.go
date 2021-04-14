package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Sid")
	want := "Hello, Sid!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
