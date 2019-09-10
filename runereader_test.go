package utf8n_test

import (
	"github.com/reiver/go-utf8n"

	"io"
	"strings"

	"testing"
)

func TestRuneReader(t *testing.T) {

	const hbegin = "\x1b[93;41m"
	const hend   = "\x1b[0m"

	tests := []struct{
		Text     string
		Expected string
	}{
		{
			Text:     "",
			Expected: "",
		},



		{
			Text:     " ",
			Expected: " ",
		},
		{
			Text:     "  ",
			Expected: "  ",
		},
		{
			Text:     "   ",
			Expected: "   ",
		},
		{
			Text:     "    ",
			Expected: "    ",
		},
		{
			Text:     "     ",
			Expected: "     ",
		},



		{
			Text:     "\t",
			Expected: "\t",
		},
		{
			Text:     "\t\t",
			Expected: "\t\t",
		},
		{
			Text:     "\t\t\t",
			Expected: "\t\t\t",
		},
		{
			Text:     "\t\t\t\t",
			Expected: "\t\t\t\t",
		},
		{
			Text:     "\t\t\t\t\t",
			Expected: "\t\t\t\t\t",
		},



		{
			Text:     "apple",
			Expected: "apple",
		},
		{
			Text:     "BANANA",
			Expected: "BANANA",
		},
		{
			Text:     "Cherry",
			Expected: "Cherry",
		},
		{
			Text:     "dATE",
			Expected: "dATE",
		},



		{
			Text:     "👾",
			Expected: "👾",
                },
		{
			Text:     "👾👻",
			Expected: "👾👻",
                },



		{
			Text:     "\r\n",
			Expected: "\u2028", // line separator
                },
		{
			Text:     "\n",
			Expected: "\u2028", // line separator
                },
		{
			Text:     "\r",
			Expected: "\u2028", // line separator
                },
		{
			Text:     "\u0085", // next line
			Expected: "\u2028", // line separator
                },
		{
			Text:     "\u2028", // line separator
			Expected: "\u2028", // line separator
                },



		{
			Text:     "\r\n\r\n",
			Expected: "\u2029", // paragraph separator
                },
		{
			Text:     "\n\n",
			Expected: "\u2029", // paragraph separator
                },
		{
			Text:     "\r\r",
			Expected: "\u2029", // paragraph separator
                },
		{
			Text:     "\u0085\u0085", // next line
			Expected: "\u2029", //paragraph separator
                },
		{
			Text:     "\u2028\u2028", // line separator
			Expected: "\u2029", // paragraph separator
                },
		{
			Text:     "\u2029", // paragraph separator
			Expected: "\u2029", // paragraph separator
                },



		{
			Text:     "\r\n"  +"wow!",
			Expected: "\u2028"+"wow!", // line separator
                },
		{
			Text:     "\n"    +"wow!",
			Expected: "\u2028"+"wow!", // line separator
                },
		{
			Text:     "\r"    +"wow!",
			Expected: "\u2028"+"wow!", // line separator
                },
		{
			Text:     "\u0085"+"wow!", // next line
			Expected: "\u2028"+"wow!", // line separator
                },
		{
			Text:     "\u2028"+"wow!", // line separator
			Expected: "\u2028"+"wow!", // line separator
                },



		{
			Text:     "\r\n\r\n"+"wow!",
			Expected: "\u2029"  +"wow!", // paragraph separator
                },
		{
			Text:     "\n\n"  +"wow!",
			Expected: "\u2029"+"wow!", // paragraph separator
                },
		{
			Text:     "\r\r"  +"wow!",
			Expected: "\u2029"+"wow!", // paragraph separator
                },
		{
			Text:     "\u0085\u0085"+"wow!", // next line
			Expected: "\u2029"      +"wow!", //paragraph separator
                },
		{
			Text:     "\u2028\u2028"+"wow!", // line separator
			Expected: "\u2029"      +"wow!", // paragraph separator
                },



		{
			Text:     "Hello world!"+"\u2028"+"wow!"+"\u0085"+"apple"+"\r"    +"BANANA"+"\n"    +"Cherry"+"\r\n"  +"dATE"+"\r\n\r\n",
			Expected: "Hello world!"+"\u2028"+"wow!"+"\u2028"+"apple"+"\u2028"+"BANANA"+"\u2028"+"Cherry"+"\u2028"+"dATE"+"\u2029",
                },
	}

	TestLoop: for testNumber, test := range tests {

		{
			var readSeeker io.ReadSeeker = strings.NewReader(test.Text)

			var runeReader io.RuneReader = utf8n.RuneReader(readSeeker)

			var dst strings.Builder

			var iterationNumber int = -1
			InnerLoop: for {
				iterationNumber++

				//if iterationNumber > len(test.Text) {
				if iterationNumber > 10 + len(test.Text) {
					t.Errorf("For test #%d and iteration #%d, %siterated way too many times!%s", testNumber, iterationNumber, hbegin, hend)

					t.Logf("TEXT: %q", test.Text)
					t.Logf("DST: %q", dst.String())
					continue TestLoop
				}

				r, size, err := runeReader.ReadRune()
				if nil != err && io.EOF == err {
					break InnerLoop
				}
				if nil != err {
					t.Errorf("For test #%d and iteration #%d, %sdid not expect an error, but actually got one.%s", testNumber, iterationNumber, hbegin, hend)

					t.Logf("TEXT: %q", test.Text)
					t.Logf("DST: %q", dst.String())

					t.Logf("ERROR TYPE: %T", err)
					t.Logf("ERROR: %q", err)
					continue TestLoop
				}
				if expected, actual := len(string(r)), size; expected != actual {
					t.Errorf("For test #%d and iteration #%d, %sactual returned size does not match expected UTF-8 size of rune.%s", testNumber, iterationNumber, hbegin, hend)

					t.Logf("TEXT: %q", test.Text)
					t.Logf("DST: %q", dst.String())

					t.Logf("RUNE: %q (%d)", string(r), r)

					t.Errorf("EXPECTED: %d", expected)
					t.Errorf("ACTUAL:   %d", actual)
					continue TestLoop
				}

				dst.WriteRune(r)
			}

			if expected, actual := test.Expected, dst.String(); expected != actual {
				t.Errorf("For test #%d, what was actually read was not what was expected.", testNumber)

				t.Logf("TEXT: %q", test.Text)
				t.Logf("DST: %q", dst.String())

				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL: %q", actual)
				continue
			}
		}

	}
}
