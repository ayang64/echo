package echo

import (
	"io"
	"strings"
	"testing"
)

func TestExpandQuotedString(t *testing.T) {
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
		"single-quote": { // should read an easy to predict number of bytes
			input:         strings.NewReader(`'helloworld'`),
			expectedbytes: 12,
			shoulderror:   true,
		},
		"double-quote": { // should read an easy to predict number of bytes
			input:         strings.NewReader(`"helloworld"`),
			expectedbytes: 12,
			shoulderror:   true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			nbytes, _, err := readatom(test.input)
			if got, expected := nbytes, test.expectedbytes; got != expected {
				t.Fatalf("readatom(%v) returned %d bytes; expected %d bytes", test.input, got, expected)
			}

			if got, expected := test.shoulderror, err != nil; got != expected {
				s := func(e bool) string {
					if e {
						return "non-nil"
					}
					return "nil"
				}
				t.Fatalf("whitespace(%v) returned a %s error value; expected %s", test.input, s(got), s(expected))
			}
		})
	}
}

func TestExpandAtom(t *testing.T) {
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
		"chars": { // should read an easy to predict number of bytes
			input:         strings.NewReader(`helloworld`),
			expectedbytes: 10,
			shoulderror:   true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			nbytes, _, err := readatom(test.input)
			if got, expected := nbytes, test.expectedbytes; got != expected {
				t.Fatalf("readatom(%v) returned %d bytes; expected %d bytes", test.input, got, expected)
			}

			if got, expected := test.shoulderror, err != nil; got != expected {
				s := func(e bool) string {
					if e {
						return "non-nil"
					}
					return "nil"
				}
				t.Fatalf("whitespace(%v) returned a %s error value; expected %s", test.input, s(got), s(expected))
			}
		})
	}
}

func TestExpandWhiteSpace(t *testing.T) {
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
			nbytes, _, err := whitespace(test.input)
			if got, expected := nbytes, test.expectedbytes; got != expected {
				t.Fatalf("whitespace(%v) returned %d bytes; expected %d bytes", test.input, got, expected)
			}

			if got, expected := test.shoulderror, err != nil; got != expected {
				s := func(e bool) string {
					if e {
						return "non-nil"
					}
					return "nil"
				}
				t.Fatalf("whitespace(%v) returned a %s error value; expected %s", test.input, s(got), s(expected))
			}
		})
	}
}
