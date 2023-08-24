package main

import (
	"fmt"
	"testing"
)

type TestCase struct {
	expected string
	actual   string
}

func TestUnpack(t *testing.T) {
	tests := []TestCase{
		TestCase{expected: "aaaabccddddde", actual: "a4bc2d5e"},
		TestCase{expected: "ggddddde", actual: "g2d5e"},
	}

	for _, test := range tests {
		actual := unpackString(test.actual)
		if actual != test.expected {
			t.Errorf("test failed, expected: %s, got: %s", test.expected, actual)
		} else {
			fmt.Println("test passed")
		}
	}
}
