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

	var buffer strings.Builder

	for {
		r, _, err := runeScanner.ReadRune()
		if nil != err && io.EOF == err {
			break
		}
		if nil != err {
			fmt.Printf("ERROR: problem getting next rune: %s\n", err)
			return
		}

		switch r {
		case '\t':
			fmt.Printf("%q\n", buffer.String())
			buffer.Reset()

			fmt.Printf("%q (tab)\n", string(r))

		case ' ':
			fmt.Printf("%q\n", buffer.String())
			buffer.Reset()

			fmt.Printf("%q (space)\n", string(r))

		case '\u2028':
			fmt.Printf("%q\n", buffer.String())
			buffer.Reset()

			fmt.Printf("%q (line separator)\n", string(r))

		case '\u2029':
			fmt.Printf("%q\n", buffer.String())
			buffer.Reset()

			fmt.Printf("%q (paragraph separator)\n", string(r))

		default:
			buffer.WriteRune(r)
		}
	}
	if 0 < buffer.Len() {
		fmt.Printf("%q\n", buffer.String())
		buffer.Reset()
	}

// Output:
// "Hello"
// " " (space)
// "world!"
// "\u2029" (paragraph separator)
// "Khodafez."
// "\u2029" (paragraph separator)
// "apple"
// "\u2028" (line separator)
// "BANANA"
// "\u2028" (line separator)
// "Cherry"
// "\u2028" (line separator)
// "dATE"
// "\u2028" (line separator)
}
