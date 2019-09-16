
/*
Package md supports markdown files
*/
package md

import (
	"io/ioutil"
	"strings"
	"unicode"
)

var ch rune

/*
GetSource returns code sections eclosed in triple backticks.
*/
func GetSource(mdfile string) (string, error) {
	inbuf, err := ioutil.ReadFile(mdfile)
	if err != nil {
		return "", err
	}
	in, out := strings.NewReader(string(inbuf)), new(strings.Builder)
	ch = next(in)
	for in.Len() > 0 {
		switch ch {
		case '\u0060':
			out.WriteString(space(ch))
			ch = next(in)
			if ch == '\u0060' {
				out.WriteString(space(ch))
				ch = next(in)
				if ch == '\u0060' {
					out.WriteString(space(ch))
					writeSpec(in, out)
					ch = next(in)
				}
			}
		default:
			out.WriteString(space(ch))
			ch = next(in)
		}
	}
	return out.String(), nil
}

func space(ch rune) string {
	if unicode.IsSpace(ch) {
		return string(ch)
	}
	return " "
}

func writeSpec(in *strings.Reader, out *strings.Builder) {
	ch = next(in)
	for in.Len() > 0 {
		switch {
		case ch == '\u0060':
			ch = next(in)
			if ch == '\u0060' {
				ch = next(in)
				if ch == '\u0060' {
					out.WriteString("   ")
					return
				}
				out.WriteString("\u0060\u0060")
			} else {
				out.WriteString("\u0060")
			}
			return
		default:
			out.WriteRune(ch)
			ch = next(in)
		}
	}
}

func next(in *strings.Reader) rune {
	if in.Len() <= 0 {
		return -1
	}
	ch, _, err := in.ReadRune()
	if err != nil {
		panic(err)
	}
	return ch
}
