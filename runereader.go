package utf8n

import (
	"fmt"
	"io"
	"unicode/utf8"
)

func RuneScanner(readSeeker io.ReadSeeker) io.RuneScanner {
	return &internalRuneScanner {
		readSeeker:readSeeker,
	}
}

type internalRuneScanner struct {
	readSeeker io.ReadSeeker
	unread bool
	previous rune
}

func (receiver *internalRuneScanner) yield(value rune) (r rune, size int, err error) {
	receiver.previous = value
	return value, utf8.RuneLen(value), nil
}

func (receiver *internalRuneScanner) ReadRune() (r rune, size int, err error) {

	if receiver.unread {
		receiver.unread = false

		r = receiver.previous

		return r, utf8.RuneLen(r), nil
	}

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

	return receiver.yield(r)
}

func (receiver *internalRuneScanner) UnreadRune() error {
	receiver.unread = true
	return nil
}
