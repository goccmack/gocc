package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	fe2 "github.com/maxcalandrelli/gocc/internal/fe2"
)

func showResult(r interface{}, e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "parsing returned the following error: %s\n", e.Error())
	} else {
		fmt.Printf("%#v\n", r)
	}
}

var (
	File string
	Text string
)

func main() {
	flag.StringVar(&File, "file", "", "parse also text in file")
	flag.StringVar(&Text, "text", "", "parse also text given with flag")
	flag.Parse()
	if Text > "" {
		showResult(fe2.ParseText(Text))
	}
	if File > "" {
		showResult(fe2.ParseFile(File))
	}
	if str := strings.Join(os.Args[1:], " "); str > "" {
		showResult(fe2.ParseText(str))
	}
}
