package main

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	current := getTime()[0:16]
	expected := time.Now().String()[0:16]

	if current != expected {
		t.Errorf("test failed expected: %s, got: %s", expected, current)
	}
}
