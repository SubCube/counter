package main

import "testing"

func TestCountWords(t *testing.T) {

	type testCase struct {
		name        string
		input       string
		expectation int
	}

	testCases := []testCase{
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
		result := CountWords([]byte(tc.input))

		t.Run(tc.name, func(t *testing.T) {
			if result != tc.expectation {
				t.Logf("expected %d result %d", tc.expectation, result)
				t.Fail()
			}
		})
	}

}

//  go test .
// go test ./...
