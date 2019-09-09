package utf8n

import (
	"unicode/utf8"
)

// decodeRunePS only DIRECTLY does does the PS (paragraph separator) transformations,
// but also does the LS (line separator) transformations INDIRECTLY.
//
// The LS (line separator) transformations are actually done by decodeRuneLS.
//
// However, this function calls decodeRuneLS.
// So it ends up doing the LS (line separator) transformations too.
//
// So, by calling decodeRuneLS the following transformations are first done:...
//
//	CR LF ⇒ LS
//
//	LF    ⇒ LS
//
//	CR    ⇒ LS
//
//	NEL   ⇒ LS
//
// And then, after that, it does the following transformations:/..
//
//	LS LS ⇒ PS
func decodeRunePS(p []byte) (r rune, size int) {

	r, size = decodeRuneLS(p)
	if utf8.RuneError == r {
		return RuneError, size
	}

	switch r {
	case LS:
		// Nothing here.
	default:
		return r, size
	}

	{
		remaining := len(p) - size
		if 0 == remaining {
			return r, size
		}
	}

	{
		r2, size2 := decodeRuneLS(p[size:])
		if utf8.RuneError == r2 {
			return RuneError, size+size2
		}

		if LS == r && LS == r2 {
			return PS, size+size2
		}
	}

	return r, size
}
