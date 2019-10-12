package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	gocc2 "github.com/maxcalandrelli/gocc/internal/frontend/reparsed"
)

func showResult(r interface{}, e error, l int) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "parsing returned the following error: %s\n", e.Error())
	} else {
		fmt.Printf("r=%#v, %d bytes\n", r, l)
	}
}

var (
	File    string
	Text    string
	Longest bool
)

func parse(longest bool, lex *gocc2.Lexer) (res interface{}, err error, ptl int) {
	if longest {
		return gocc2.NewParser().ParseLongestPrefix(lex)
	} else {
		return gocc2.NewParser().Parse(lex)
	}
	return
}

func main() {
	flag.StringVar(&File, "file", "", "parse also text in file")
	flag.StringVar(&Text, "text", "", "parse also text given with flag")
	flag.BoolVar(&Longest, "longest", false, "parse longest possible part")
	flag.Parse()
	if Text > "" {
		showResult(parse(Longest, gocc2.NewLexerString(Text)))
	}
	if File > "" {
		l, e := gocc2.NewLexerFile(File)
		if e != nil {
			panic(e)
		}
		showResult(parse(Longest, l))
	}
	if str := strings.Join(flag.Args(), " "); str > "" {
		showResult(parse(Longest, gocc2.NewLexerString(str)))
	}
}
