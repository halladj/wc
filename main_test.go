package main

import (
	"bytes"
	"testing"
)

func TestCountWord(t *testing.T) {

	b := bytes.NewBufferString("hey ronaldo im messi\n")

	expected := 4

	res := count(b, false)

	if res != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, res)
	}
}

func TestCountLines(t *testing.T) {

	b := bytes.NewBufferString("hey ronaldo im messi\nyou are\nmy goat")

	expected := 3

	res := count(b, true)

	if res != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, res)
	}
}
