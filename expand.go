package echo

import (
	"io"
	"unicode"
	"fmt"
	"strings"
)


// eatwhitespace returns the number of bytes it consume and any errors it
// encounters along the way.
//
func eatwhitespace(rs io.RuneScanner) (int, error) {
	if rs == nil {
		return 0, fmt.Errorf("nil scanner")
	}

	var sum int
	for {
		r, nbytes, err := rs.ReadRune()

		// if an error occured, return immediately.
		if err != nil {
			return sum, err
		}

		// if we didn't get an error, ensure that the rune we've read is is a
		// whitespace.
		if !unicode.IsSpace(r) {
			// rune is not a space. put it back into the input stream and break out
			// of loop.
			rs.UnreadRune()
			break
		}

		// we've foun a white space characater at this point.  add it to our bytes
		// read.
		sum += nbytes
	}
	return sum, nil
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
