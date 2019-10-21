package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	calc "github.com/maxcalandrelli/gocc/example/calc/calc.grammar/calc"
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

func parse(longest bool, lex *calc.Lexer) (res interface{}, err error, parsed []byte) {
	if longest {
		return calc.NewParser().ParseLongestPrefix(lex)
	} else {
		res, err = calc.NewParser().Parse(lex)
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
		showResult(parse(Longest, calc.NewLexerString(Text)))
	}
	if File > "" {
		l, e := calc.NewLexerFile(File)
		if e != nil {
			panic(e)
		}
		showResult(parse(Longest, l))
	}
	if str := strings.Join(flag.Args(), " "); str > "" {
		showResult(parse(Longest, calc.NewLexerString(str)))
	}
}
