/*
Package terminput provides terminal keyboard input for interactive command-line tools.

Note that this package does not enable raw mode, use github.com/pkg/term or similar.
*/
package terminput

import (
	"errors"
	"io"
	"unicode/utf8"
)

// Read terminal input from the given reader.
func Read(r io.Reader) (*KeyboardInput, error) {
	var buf [256]byte

	n, err := r.Read(buf[:])
	if err != nil {
		return nil, err
	}

	c, _ := utf8.DecodeRune(buf[:])
	if c == utf8.RuneError {
		return nil, errors.New("invalid rune")
	}

	// control characters
	if n == 1 && Key(c) <= KeyUS {
		return newKeyboardInput(Key(c), c, ModNone), nil
	}

	// sequences
	key, ok := keySequences[string(buf[:n])]
	if ok {
		var mod Mod

		switch key {
		case keyShiftLeft:
			key = KeyLeft
			mod |= ModShift
		case keyShiftRight:
			key = KeyRight
			mod |= ModShift
		case keyAltLeft:
			key = KeyLeft
			mod |= ModAlt
		case keyAltRight:
			key = KeyRight
			mod |= ModAlt
		}

		return newKeyboardInput(key, ' ', mod), nil
	}

	return newKeyboardInput(KeyRune, c, ModNone), nil
}
