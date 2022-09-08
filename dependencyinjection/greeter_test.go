package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "bob")

	got := buffer.String()
	want := "hi bob"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}