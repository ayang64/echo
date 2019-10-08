package echo

import (
	"io"
	"unicode"
	"fmt"
	"strings"
)

func quotedstring(rs io.RuneScanner) (int, string, error) {
	var qtype rune
	var nrunes int
	var closed bool

	start := func() bool {
		return nrunes == 1
	}

	qstring := func(r rune) bool {
		nrunes++
		// if we've closed the string, return false
		if closed {
			return false
		}

		// if we're reading the first rune, check to see if it is a quote
		// character.
		if start() {
			switch r {
			case '\'', '"':
				qtype = r
				return true
			default:
				return false
			}
		}

		// if we've read a closing quote, close the string
		if r == qtype {
			closed = true
		}

		return true
	}

	return match(rs, qstring)
}

func atom(rs io.RuneScanner) (int, string, error) {
	atom := func(r rune) bool {
		return !unicode.IsSpace(r)
	}
	return match(rs, atom)
}

// whitespace returns the number of bytes it consume and any errors it
// encounters along the way.
//
func whitespace(rs io.RuneScanner) (int, string, error) {
	return match(rs, unicode.IsSpace)
}

func match(rs io.RuneScanner, matchfunc func(rune) bool) (int, string, error) {
	if rs == nil {
		return 0, "", fmt.Errorf("nil scanner")
	}

	sb := strings.Builder{}

	var sum int
	for {
		r, nbytes, err := rs.ReadRune()

		// if an error occured, return immediately.
		if err != nil {
			return sum, sb.String(), err
		}

		// if we didn't get an error, ensure that the rune we've read is is a
		// whitespace.
		if !matchfunc(r) {
			// rune is not a space. put it back into the input stream and break out
			// of loop.
			rs.UnreadRune()
			break
		}

		// we've foun a white space characater at this point.  add it to our bytes
		// read.
		sum += nbytes
		sb.WriteRune(r)
	}
	return sum, sb.String(), nil
}

func (e *echo) expand(rs io.RuneScanner) ([]string, error) {
	return nil, nil
}

// Expand splits a string into tokens in a similar if simplified way to the
// shell.
//
// Expand() currently identifies two tokens:
//
//   atom - a contiguous group of non-white space characters
//
//   qstring - a quoted string that begins and ends with either double quotes
//   (") or single quotes (').
//
// This allows us to quote the values passed to echo.Run() like:
//
//   e.Run(`rm -fr "/tmp/File With Spaces In the Name"`)
//
// Expand actually doesn't expand shell variables or evaluate any other
// shell expressions.
//
// It is named after the expand() function in bourne shell found under FreeBSD
// which does similar splitting.
//
// A more apropriate name would probably be Split() but it is already taken.
//
func (e *echo) Expand(rs io.RuneScanner) ([]string, error) {
	return nil, nil
}


func (e *echo) ExpandString(s string) ([]string, error) {
	return e.Expand(strings.NewReader(s))
}
