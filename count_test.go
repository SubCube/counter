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

func TestCountLines(t *testing.T) {

	testCases := []struct {
		name        string
		input       string
		expectation int
	}{
		{
			name:        "simple 5 words, 1 new line",
			input:       "one two three four five\n",
			expectation: 1,
		},
		{
			name:        "empty file",
			input:       "",
			expectation: 0,
		},
		{
			name:        "no new line",
			input:       "one two three four five",
			expectation: 0,
		},
		{
			name:        "no new line at end",
			input:       "one two three four five\n six",
			expectation: 1,
		},
		{
			name:        "4 new lines only",
			input:       "\n\n\n\n",
			expectation: 4,
		},
		{
			name:        "multi word new lines",
			input:       "one\ntwo\nthree\nfour\nfive\n",
			expectation: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)

			numLines := counter.CountLines(r)

			if numLines != tc.expectation {
				t.Logf("expected %d result %d", tc.expectation, numLines)
				t.Fail()
			}
		})
	}

}

// go test .
// go test ./...
// go test . -v
