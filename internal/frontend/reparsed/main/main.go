package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	reparsed "github.com/maxcalandrelli/gocc/internal/frontend/reparsed"
)

func showResult(r interface{}, e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "parsing returned the following error: %s\n", e.Error())
	} else {
		fmt.Printf("%#v\n", r)
	}
}

var (
	File    string
	Text    string
	Longest bool
)

func main() {
	flag.StringVar(&File, "file", "", "parse also text in file")
	flag.StringVar(&Text, "text", "", "parse also text given with flag")
	flag.BoolVar(&Longest, "longest", false, "parse longest possible part")
	flag.Parse()
	if Text > "" {
		showResult(reparsed.ParseText(Text))
	}
	if File > "" {
		showResult(reparsed.ParseFile(File))
	}
	if str := strings.Join(os.Args[1:], " "); str > "" {
		showResult(reparsed.ParseText(str))
	}
}
