package main

import "testing"

func TestShell(t *testing.T) {
	type testCase struct {
		actual   string
		expected string
	}
	testCases := []testCase{
		{
			actual:   pwd(),
			expected: "/Users/seleroad/GolandProjects/L2/dev08",
		},
	}

	for _, test := range testCases {
		if test.actual != test.expected {
			t.Errorf("expected: %s, got: %s", test.expected, test.actual)
		}
	}

}
