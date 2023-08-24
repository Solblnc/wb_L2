package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		arg []string
	}{
		{arg: []string{"run", "task.go", "-k", "3", "example.txt"}},
		{arg: []string{"run", "task.go", "-r", "example.txt"}},
		{arg: []string{"run", "task.go", "-u", "example.txt"}},
	}

	file, err := os.Open("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	file.Close()

	for _, v := range testCases {
		t1, _ := exec.Command("go", v.arg...).Output()
		t2, _ := exec.Command("sort", v.arg[2:]...).Output()

		if string(t1) != string(t2) {
			t.Fatalf("\nExpected: %s \nActual: %s", t2, t1)
		}

	}

}
