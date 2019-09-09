package utf8n

import (
	"unicode/utf8"
)

// decodeRuneLS only does the LS (line separator) transformations.
//
// (The PS (paragraph separator) transformations are done by decodeRunePS
// which calls this function.)
//
// I.e.,...
//
//	CR LF ⇒ LS
//
//	LF    ⇒ LS
//
//	CR    ⇒ LS
//
//	NEL   ⇒ LS
func decodeRuneLS(p []byte) (r rune, size int) {

	r, size = utf8.DecodeRune(p)
	if utf8.RuneError == r {
		return RuneError, size
	}

	switch r {
	case LF, CR, NEL:
		// Nothing here.
	default:
		return r, size
	}

	{
		remaining := len(p) - size

		if 0 == remaining {
			switch r {
			case LF, CR, NEL:
				return LS, size
			default:
				return r, size
			}
		}
	}

	{
		r2, size2 := utf8.DecodeRune(p[size:])
		if utf8.RuneError == r2 {
			return RuneError, size+size2
		}

		if CR == r && LF == r2 {
			return LS, size+size2
		}
	}

	switch r {
	case LF, CR, NEL:
		return LS, size
	default:
		return r, size
	}
}
