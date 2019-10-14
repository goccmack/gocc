package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	bools "github.com/maxcalandrelli/gocc/example/bools/bools.grammar/bools"
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

func parse(longest bool, lex *bools.Lexer) (res interface{}, err error, parsed []byte) {
	if longest {
		return bools.NewParser().ParseLongestPrefix(lex)
	} else {
		res, err = bools.NewParser().Parse(lex)
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
		showResult(parse(Longest, bools.NewLexerString(Text)))
	}
	if File > "" {
		l, e := bools.NewLexerFile(File)
		if e != nil {
			panic(e)
		}
		showResult(parse(Longest, l))
	}
	if str := strings.Join(flag.Args(), " "); str > "" {
		showResult(parse(Longest, bools.NewLexerString(str)))
	}
}
