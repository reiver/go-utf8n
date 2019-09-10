package utf8n

import (
	"errors"
)

var (
	errNilReadSeeker = errors.New("utf8n: Nil io.ReadSeeker")
)
