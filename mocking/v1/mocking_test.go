package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := "3\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
