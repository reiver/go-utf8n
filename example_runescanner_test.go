package utf8n_test

import (
	"github.com/reiver/go-utf8n"

	"fmt"
	"io"
	"strings"
)

func ExampleRuneScanner() {

	var text =
`Hello world!

Khodafez.

apple
BANANA
Cherry
dATE
`

	var readSeeker io.ReadSeeker = strings.NewReader(text)

	var runeScanner io.RuneScanner = utf8n.RuneScanner(readSeeker)

	var r rune
	var err error

	for i:=0; i<22; i++ {
		r, _, err = runeScanner.ReadRune()
		if nil != err {
			fmt.Printf("ERROR: could not read another rune: %s", err)
			return
		}

		fmt.Printf("%q\n", string(r))
	}

	fmt.Printf("=====])> UNREAD! %q\n", string(r))

	err = runeScanner.UnreadRune()
	if nil != err {
		fmt.Printf("ERROR: could not unread rune: %s", err)
		return
	}

	for i:=0; i<27; i++ {
		r, _, err = runeScanner.ReadRune()
		if nil != err {
			fmt.Printf("ERROR: could not read another rune: %s", err)
			return
		}

		fmt.Printf("%q\n", string(r))
	}

// Output:
// "H"
// "e"
// "l"
// "l"
// "o"
// " "
// "w"
// "o"
// "r"
// "l"
// "d"
// "!"
// "\u2029"
// "K"
// "h"
// "o"
// "d"
// "a"
// "f"
// "e"
// "z"
// "."
// =====])> UNREAD! "."
// "."
// "\u2029"
// "a"
// "p"
// "p"
// "l"
// "e"
// "\u2028"
// "B"
// "A"
// "N"
// "A"
// "N"
// "A"
// "\u2028"
// "C"
// "h"
// "e"
// "r"
// "r"
// "y"
// "\u2028"
// "d"
// "A"
// "T"
// "E"
// "\u2028"
}
