package utf8n

import (
	"fmt"
	"io"
	"unicode/utf8"
)

func RuneReader(readSeeker io.ReadSeeker) io.RuneReader {
	return internalRuneReader {
		readSeeker:readSeeker,
	}
}

type internalRuneReader struct {
	readSeeker io.ReadSeeker
}

func (receiver internalRuneReader) ReadRune() (r rune, size int, err error) {

	var readSeeker io.ReadSeeker = receiver.readSeeker
	if nil == readSeeker {
		return RuneError, 0, errNilReadSeeker
	}

	// This is UTF8Max*2 (rather than just UTF8Max) because of "\r\n"
	var buffer [UTF8Max * 2]byte
	var src []byte = buffer[:]
	var numRead int
	{
		numRead, err = readSeeker.Read(src)
		if nil != err {
			return RuneError, numRead, err
		}
		if numRead > len(buffer) {
			return RuneError, numRead, fmt.Errorf("utf8n: Internal Error: number of bytes read (%d) is larger than the buffer (%d)", numRead, len(buffer))
		}
		src = src[:numRead]
	}

	r, size = DecodeRune(src)

	{
		difference := numRead - size

		if 0 < difference {
			_, err := readSeeker.Seek(int64(difference * -1), io.SeekCurrent)
			if nil != err {
				return r, -1, err
			}
		}
	}

	return r, utf8.RuneLen(r), nil
}
