package utf8n

import (
	"unicode/utf8"

	"testing"
)

func TestDecodeRunePS(t *testing.T) {

	tests := []struct{
		Bytes []byte
		ExpectedRune rune
		ExpectedSize int
	}{
		{
			Bytes: []byte("apple"),
			ExpectedRune: 'a',
			ExpectedSize: utf8.RuneLen('a'),
		},
		{
			Bytes: []byte("BANANA"),
			ExpectedRune: 'B',
			ExpectedSize: utf8.RuneLen('B'),
		},
		{
			Bytes: []byte("Cherry"),
			ExpectedRune: 'C',
			ExpectedSize: utf8.RuneLen('C'),
		},
		{
			Bytes: []byte("dATE"),
			ExpectedRune: 'd',
			ExpectedSize: utf8.RuneLen('d'),
		},



		{
			Bytes: []byte("Hello world"),
			ExpectedRune: 'H',
			ExpectedSize: utf8.RuneLen('H'),
		},



		{
			Bytes: []byte("👾"),
			ExpectedRune: '👾',
			ExpectedSize: utf8.RuneLen('👾'),
		},
		{
			Bytes: []byte("👾👻"),
			ExpectedRune: '👾',
			ExpectedSize: utf8.RuneLen('👾'),
		},





		{
			Bytes: []byte(string(CR)+string(LF)),
			ExpectedRune: LS,
			ExpectedSize: utf8.RuneLen(CR)+utf8.RuneLen(LF),
		},
		{
			Bytes: []byte(string(LF)),
			ExpectedRune: LS,
			ExpectedSize: utf8.RuneLen(LF),
		},
		{
			Bytes: []byte(string(CR)),
			ExpectedRune: LS,
			ExpectedSize: utf8.RuneLen(CR),
		},
		{
			Bytes: []byte(string(NEL)),
			ExpectedRune: LS,
			ExpectedSize: utf8.RuneLen(NEL),
		},
		{
			Bytes: []byte(string(LS)),
			ExpectedRune: LS,
			ExpectedSize: utf8.RuneLen(LS),
		},



		{
			Bytes: []byte(string(CR)+string(LF)+"wow!"),
			ExpectedRune: LS,
			ExpectedSize: utf8.RuneLen(CR)+utf8.RuneLen(LF),
		},
		{
			Bytes: []byte(string(LF)+"wow!"),
			ExpectedRune: LS,
			ExpectedSize: utf8.RuneLen(LF),
		},
		{
			Bytes: []byte(string(CR)+"wow!"),
			ExpectedRune: LS,
			ExpectedSize: utf8.RuneLen(CR),
		},
		{
			Bytes: []byte(string(NEL)+"wow!"),
			ExpectedRune: LS,
			ExpectedSize: utf8.RuneLen(NEL),
		},
		{
			Bytes: []byte(string(LS)+"wow!"),
			ExpectedRune: LS,
			ExpectedSize: utf8.RuneLen(LS),
		},









		{
			Bytes: []byte(string(CR)+string(LF)+string(CR)+string(LF)),
			ExpectedRune: PS,
			ExpectedSize: utf8.RuneLen(CR)+utf8.RuneLen(LF)+utf8.RuneLen(CR)+utf8.RuneLen(LF),
		},
		{
			Bytes: []byte(string(LF)+string(LF)),
			ExpectedRune: PS,
			ExpectedSize: utf8.RuneLen(LF)+utf8.RuneLen(LF),
		},
		{
			Bytes: []byte(string(CR)+string(CR)),
			ExpectedRune: PS,
			ExpectedSize: utf8.RuneLen(CR)+utf8.RuneLen(CR),
		},
		{
			Bytes: []byte(string(NEL)+string(NEL)),
			ExpectedRune: PS,
			ExpectedSize: utf8.RuneLen(NEL)+utf8.RuneLen(NEL),
		},
		{
			Bytes: []byte(string(LS)+string(LS)),
			ExpectedRune: PS,
			ExpectedSize: utf8.RuneLen(LS)+utf8.RuneLen(LS),
		},
		{
			Bytes: []byte(string(PS)),
			ExpectedRune: PS,
			ExpectedSize: utf8.RuneLen(PS),
		},



		{
			Bytes: []byte(string(CR)+string(LF)+string(CR)+string(LF)+"wow!"),
			ExpectedRune: PS,
			ExpectedSize: utf8.RuneLen(CR)+utf8.RuneLen(LF)+utf8.RuneLen(CR)+utf8.RuneLen(LF),
		},
		{
			Bytes: []byte(string(LF)+string(LF)+"wow!"),
			ExpectedRune: PS,
			ExpectedSize: utf8.RuneLen(LF)+utf8.RuneLen(LF),
		},
		{
			Bytes: []byte(string(CR)+string(CR)+"wow!"),
			ExpectedRune: PS,
			ExpectedSize: utf8.RuneLen(CR)+utf8.RuneLen(CR),
		},
		{
			Bytes: []byte(string(NEL)+string(NEL)+"wow!"),
			ExpectedRune: PS,
			ExpectedSize: utf8.RuneLen(NEL)+utf8.RuneLen(NEL),
		},
		{
			Bytes: []byte(string(LS)+string(LS)+"wow!"),
			ExpectedRune: PS,
			ExpectedSize: utf8.RuneLen(LS)+utf8.RuneLen(LS),
		},
		{
			Bytes: []byte(string(PS)+"wow!"),
			ExpectedRune: PS,
			ExpectedSize: utf8.RuneLen(PS),
		},
	}

	for testNumber, test := range tests {
		r, size := decodeRunePS(test.Bytes)
		if RuneError == r {
			t.Errorf("For test #%d, did not expect to get a rune error, but actually got one..", testNumber)
			t.Logf("STRING: %q (size = %d)", test.Bytes, len(test.Bytes))
			t.Logf("ACTUAL RUNE: %q (0x%x)", string(r), r)
			t.Logf("ACTUAL SIZE: %d", size)
			continue
		}

		if expected, actual := test.ExpectedRune, r; expected != actual {
			t.Errorf("For test #%d, the actual %sRUNE%s is not what was expected.", testNumber, "\x1b[93;41m", "\x1b[0m")
			t.Logf("STRING: %q (size = %d)", test.Bytes, len(test.Bytes))
			t.Logf("EXPECTED: %q (0x%x)", string(expected), expected)
			t.Logf("ACTUAL:   %q (0x%x)", string(actual), actual)
			continue
		}

		if expected, actual := test.ExpectedSize, size; expected != actual {
			t.Errorf("For test #%d, the actual %sSIZE%s is not what was expected.", testNumber, "\x1b[97;45m", "\x1b[0m")
			t.Logf("STRING: %q (size = %d)", test.Bytes, len(test.Bytes))
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			continue
		}
	}
}
