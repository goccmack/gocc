package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	er "github.com/maxcalandrelli/gocc/example/errorrecovery/er.grammar/er"
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

func parse(longest bool, lex *er.Lexer) (res interface{}, err error, parsed []byte) {
	if longest {
		return er.NewParser().ParseLongestPrefix(lex)
	} else {
		res, err = er.NewParser().Parse(lex)
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
		showResult(parse(Longest, er.NewLexerString(Text)))
	}
	if File > "" {
		l, e := er.NewLexerFile(File)
		if e != nil {
			panic(e)
		}
		showResult(parse(Longest, l))
	}
	if str := strings.Join(flag.Args(), " "); str > "" {
		showResult(parse(Longest, er.NewLexerString(str)))
	}
}
