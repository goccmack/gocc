package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	rr "github.com/maxcalandrelli/gocc/example/rr/rr.grammar/rr"
)

func showResult(r interface{}, e error, p []byte) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "parsing returned the following error: %s\n", e.Error())
	} else {
		if len(p) > 0 {
			fmt.Printf("r=%#v, (%s)\n", r, string(p))
		} else {
			fmt.Printf("r=%#v\n", r)
		}
	}
}

var (
	File    string
	Text    string
	Longest bool
)

func parse(longest bool, lex *rr.Lexer) (res interface{}, err error, parsed []byte) {
	if longest {
		return rr.NewParser().ParseLongestPrefix(lex)
	} else {
		res, err = rr.NewParser().Parse(lex)
		parsed = []byte{}
	}
	return
}

func main() {
	flag.StringVar(&File, "file", "", "parse also text in file")
	flag.StringVar(&Text, "text", "", "parse also text given with flag")
	flag.BoolVar(&Longest, "longest", false, "parse longest possible part")
	flag.Parse()
	if Text > "" {
		showResult(parse(Longest, rr.NewLexerString(Text)))
	}
	if File > "" {
		l, e := rr.NewLexerFile(File)
		if e != nil {
			panic(e)
		}
		showResult(parse(Longest, l))
	}
	if str := strings.Join(flag.Args(), " "); str > "" {
		showResult(parse(Longest, rr.NewLexerString(str)))
	}
}
