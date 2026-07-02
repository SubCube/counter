package main

import "testing"

func TestCountWords(t *testing.T) {
	input := "one two three four five"
	expectation := 5

	result := CountWords([]byte(input))

	if result != expectation {
		t.Logf("expected %d result %d", expectation, result)
		t.Fail()
	}

	// edge case 1
	input = ""
	expectation = 0
	result = CountWords([]byte(input))

	if result != expectation {
		t.Log("expected", expectation, "result", result)
		t.Fail()
	}

	// edge case 2
	input = " "
	expectation = 0
	result = CountWords([]byte(input))

	if result != expectation {
		t.Log("expected", expectation, "result", result)
		t.Fail()
	}
}

//  go test .
// go test ./...
