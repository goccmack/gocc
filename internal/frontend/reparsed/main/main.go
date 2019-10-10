package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	reparsed "github.com/maxcalandrelli/gocc/internal/frontend/reparsed"
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

func parse(longest bool, lex *reparsed.Lexer) (res interface{}, err error, ptl int) {
	if longest {
		return reparsed.NewParser().ParseLongestPrefix(lex)
	} else {
		return reparsed.NewParser().Parse(lex)
	}
	return
}

func main() {
	flag.StringVar(&File, "file", "", "parse also text in file")
	flag.StringVar(&Text, "text", "", "parse also text given with flag")
	flag.BoolVar(&Longest, "longest", false, "parse longest possible part")
	flag.Parse()
	if Text > "" {
		showResult(parse(Longest, reparsed.NewLexerString(Text)))
	}
	if File > "" {
		l, e := reparsed.NewLexerFile(File)
		if e != nil {
			panic(e)
		}
		showResult(parse(Longest, l))
	}
	if str := strings.Join(flag.Args(), " "); str > "" {
		showResult(parse(Longest, reparsed.NewLexerString(str)))
	}
}
