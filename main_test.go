package main

import "testing"

func TestCountWords(t *testing.T) {
	input := "one two three four five"
	expectation := 5

	result := countWords([]byte(input))

	if result != expectation {
		t.Fail()
	}
}

//  go test .
// go test ./...
