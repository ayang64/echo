package echo

import (
	"io"
	"strings"
	"testing"
)

func TestEatWhiteSpace(t *testing.T) {
	tests := map[string]struct {
		input         io.RuneScanner
		expectedbytes int
		shoulderror   bool
	}{
		"nil reader": {
			input:         nil,
			expectedbytes: 0,
			shoulderror:   true,
		},
		"empty string": {
			input:         strings.NewReader(""),
			expectedbytes: 0,
			shoulderror:   true, // should return io.EOF.  this test should check for that but i'm lazy
		},
		"8 bit string": { // should read an easy to predict number of bytes
			input:         strings.NewReader("          "),
			expectedbytes: 10,
			shoulderror:   true,
		},
		"tabs-spaces-and-newlines": { // should read an easy to predict number of bytes
			input:         strings.NewReader("  \t\n\t\t  \n\n  "),
			expectedbytes: 12,
			shoulderror:   true,
		},
	}

	t.Parallel()

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			nbytes, err := eatwhitespace(test.input)
			if got, expected := nbytes, test.expectedbytes; got != expected {
				t.Fatalf("eatwhitespace(%v) returned %d bytes; expected %d bytes", test.input, got, expected)
			}

			if got, expected := test.shoulderror, err != nil; got != expected {
				s := func(e bool) string {
					if e {
						return "non-nil"
					}
					return "nil"
				}
				t.Fatalf("eatwhitespace(%v) returned a %s error value; expected %s", test.input, s(got), s(expected))
			}
		})
	}
}
