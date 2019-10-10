package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	errorrecovery "github.com/maxcalandrelli/gocc/example/errorrecovery"
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

func parse(longest bool, lex *errorrecovery.Lexer) (res interface{}, err error, ptl int) {
	if longest {
		return errorrecovery.NewParser().ParseLongestPrefix(lex)
	} else {
		return errorrecovery.NewParser().Parse(lex)
	}
	return
}

func main() {
	flag.StringVar(&File, "file", "", "parse also text in file")
	flag.StringVar(&Text, "text", "", "parse also text given with flag")
	flag.BoolVar(&Longest, "longest", false, "parse longest possible part")
	flag.Parse()
	if Text > "" {
		showResult(parse(Longest, errorrecovery.NewLexerString(Text)))
	}
	if File > "" {
		l, e := errorrecovery.NewLexerFile(File)
		if e != nil {
			panic(e)
		}
		showResult(parse(Longest, l))
	}
	if str := strings.Join(flag.Args(), " "); str > "" {
		showResult(parse(Longest, errorrecovery.NewLexerString(str)))
	}
}
