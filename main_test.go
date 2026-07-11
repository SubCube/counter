package main_test

import (
	"strings"
	"testing"

	counter "github.com/SubCube/counter"
)

func TestCountWords(t *testing.T) {

	testCases := []struct {
		name        string
		input       string
		expectation int
	}{
		{
			name:        "5 words",
			input:       "one two three four five",
			expectation: 5,
		},
		{
			name:        "empty",
			input:       "",
			expectation: 0,
		},
		{
			name:        "space",
			input:       " ",
			expectation: 0,
		},
	}

	for _, tc := range testCases {
		r := strings.NewReader(tc.input)
		result := counter.CountWords(r)

		t.Run(tc.name, func(t *testing.T) {
			if result != tc.expectation {
				t.Logf("expected %d result %d", tc.expectation, result)
				t.Fail()
			}
		})
	}

}

// go test .
// go test ./...
// go test . -v
