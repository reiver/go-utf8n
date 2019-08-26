package utf8n

import (
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"

	"testing"
)

func TestDecodeRuneLS(t *testing.T) {

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
			ExpectedRune: LS,
			ExpectedSize: utf8.RuneLen(CR)+utf8.RuneLen(LF),
		},
		{
			Bytes: []byte(string(LF)+string(LF)),
			ExpectedRune: LS,
			ExpectedSize: utf8.RuneLen(LF),
		},
		{
			Bytes: []byte(string(CR)+string(CR)),
			ExpectedRune: LS,
			ExpectedSize: utf8.RuneLen(CR),
		},
		{
			Bytes: []byte(string(NEL)+string(NEL)),
			ExpectedRune: LS,
			ExpectedSize: utf8.RuneLen(NEL),
		},
		{
			Bytes: []byte(string(LS)+string(LS)),
			ExpectedRune: LS,
			ExpectedSize: utf8.RuneLen(LS),
		},
		{
			Bytes: []byte(string(PS)),
			ExpectedRune: PS,
			ExpectedSize: utf8.RuneLen(PS),
		},



		{
			Bytes: []byte(string(CR)+string(LF)+string(CR)+string(LF)+"wow!"),
			ExpectedRune: LS,
			ExpectedSize: utf8.RuneLen(CR)+utf8.RuneLen(LF),
		},
		{
			Bytes: []byte(string(LF)+string(LF)+"wow!"),
			ExpectedRune: LS,
			ExpectedSize: utf8.RuneLen(LF),
		},
		{
			Bytes: []byte(string(CR)+string(CR)+"wow!"),
			ExpectedRune: LS,
			ExpectedSize: utf8.RuneLen(CR),
		},
		{
			Bytes: []byte(string(NEL)+string(NEL)+"wow!"),
			ExpectedRune: LS,
			ExpectedSize: utf8.RuneLen(NEL),
		},
		{
			Bytes: []byte(string(LS)+string(LS)+"wow!"),
			ExpectedRune: LS,
			ExpectedSize: utf8.RuneLen(LS),
		},
		{
			Bytes: []byte(string(PS)+"wow!"),
			ExpectedRune: PS,
			ExpectedSize: utf8.RuneLen(PS),
		},
	}

	{
		prefixRunes := []rune{LF, CR, NEL, LS}
		suffixRunes := []rune{
			' ','!','"','#','$','%','&','\'','(',')','*','+',',','-','.','/',
			'0','1','2','3','4','5','6','7','8','9',
			':',';','<','=','>','?','@',
			'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z',
			'[','\\',']','^','_','`',
			'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z',
			'{','|','}','~',
			'۰','۱','۲','۳','۴','۵','۶','۷','۸','۹',
			'👾',
			'🙂',
		}

		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		for i:=0; i<50; i++ {

			r0 := prefixRunes[randomness.Int() % len(prefixRunes)]
			r1 := prefixRunes[randomness.Int() % len(prefixRunes)]

			test := struct{
				Bytes []byte
				ExpectedRune rune
				ExpectedSize int
			}{}

			test.Bytes = []byte(string(r0)+string(r1))
			test.ExpectedRune = LS
			switch {
			case '\r' == r0 && '\n' == r1:
				test.ExpectedSize = utf8.RuneLen(r0)+utf8.RuneLen(r1)
			default:
				test.ExpectedSize = utf8.RuneLen(r0)
			}

			if 0 == (randomness.Int() % 2) {
				var builder strings.Builder

				rr := suffixRunes[randomness.Int() % len(suffixRunes)]
				builder.WriteRune(rr)

				{
					limit := randomness.Intn(14)
					for ii:=0; ii<limit; ii++ {
						rr2 := suffixRunes[randomness.Int() % len(suffixRunes)]
						builder.WriteRune(rr2)
					}
				}

				test.Bytes = append(test.Bytes, builder.String()...)
			}

			tests = append(tests, test)
		}
	}

	for testNumber, test := range tests {
		r, size := decodeRuneLS(test.Bytes)
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
