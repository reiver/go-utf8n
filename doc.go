/*
Package utf8n implements functions and constants to support normalizing text encoded in UTF-8.

This package is similar to the Go built-in "unicode/utf8" package,
except it normalizes ‘line separator’ and ‘paragraph separator’ characters.

So that it transforms:

	CR LF ⇒ LS

	LF    ⇒ LS

	CR    ⇒ LS

	NEL   ⇒ LS

And then after (conceptually) doing that, transforms:

	LS LS ⇒ PS

The meanings of LF, CR, NEL, LS, and PS are:

	LF  = “line feed”            = U+000A = '\u000A' = '\n'

	CR  = “carriage return”      = U+000D = '\u000D' = '\r'

	NEL = “next line”            = U+0085 = '\u0085'

	LS  = “line separator”       = U+2028 = '\u2028'

	PS  = “paragraph separator”  = U+2029 = '\u2029'

The result of these transformations is that:

№1: ‘line separator’, and ‘paragraph separator’ characters are always represented by a single rune,

№2: ‘line separator’, and ‘paragraph separator’ characters are always represented by the same runes.
*/
package utf8n
