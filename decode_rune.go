package utf8n

// DecodeRune is similar to utf8.DecodeRune() in the Go built-in "unicode/utf8" package,
// except it normalizes ‘line separator’ and ‘paragraph separator’ characters.
//
// So that it transforms:
//
//	CR LF ⇒ LS
//
//	LF    ⇒ LS
//
//	CR    ⇒ LS
//
//	NEL   ⇒ LS
//
// And then after (conceptually) doing that, transforms:
//
//	LS LS ⇒ PS
//
// The meanings of LF, CR, NEL, LS, and PS are:
//
//	LF  = “line feed”            = U+000A = '\u000A' = '\n'
//
//	CR  = “carriage return”      = U+000D = '\u000D' = '\r'
//
//	NEL = “next line”            = U+0085 = '\u0085'
//
//	LS  = “line separator”       = U+2028 = '\u2028'
//
//	PS  = “paragraph separator”  = U+2029 = '\u2029'
//
// The returned ‘size’ is the pre-transformed number of bytes read from ‘p’.
//
// That way you can do stuff such as:
//
//	p = p[size:]
//
// And the returned ‘r’ is the transformed rune.
//
// For example:
//
//	var utf8Bytes []byte = []byte("This is the 1st line\r\nThis is the second line\r\n\r\nThis is the 2nd paragraph.")
//	
//	var p []byte = utf8Bytes
//	
//	var builder strings.Builder // <--- We will put out result here.
//	
//	for {
//		r, size := utf8n.DecodeRune(p)
//		
//		if utf8n.RuneError == r && 0 == size { // Nothing more to decode.
//		
//			break // <-------------- We get out of the loop with this!
//		}
//		if utf8n.RuneError == r { // An actual error.
//			fmt.Println("ERROR: invalid UTF-8")
//			return
//		}
//		
//		p = p[size:] // <---- We skip past what we just decoded.
//		
//		builder.WriteRune(r)
//	}
//	
//	fmt.Printf("RESULT: %s\n", builder.String())
func DecodeRune(p []byte) (r rune, size int) {
	return decodeRunePS(p)
}
