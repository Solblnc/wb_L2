package main

import (
	"strings"
	"testing"
)

func TestSet(t *testing.T) {
	testCase := []string{"тяпка", "пятак", "пятка", "клоун", "локун"}
	actual := set(testCase)
	expected := make(map[string][]string)
	expected["тяпка"] = []string{"тяпка", "пятак", "пятка"}
	expected["клоун"] = []string{"клоун", "локун"}

	if len(actual) != len(expected) {
		t.Fatal("invalid size")
	}

	for key, value := range expected {
		valueExpected, exist := expected[key]
		if !exist {
			t.Fatal("key does not exist")
		}

		for i := 0; i < len(valueExpected); i++ {
			if strings.Compare(value[i], valueExpected[i]) != 0 {
				t.Fatal("strings are different")
			}
		}
	}

}
