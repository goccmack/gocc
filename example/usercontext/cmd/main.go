package main

import (
	"github.com/goccmack/gocc/example/usercontext/ast"
	"github.com/goccmack/gocc/example/usercontext/errors"
	"github.com/goccmack/gocc/example/usercontext/lexer"
	"github.com/goccmack/gocc/example/usercontext/parser"

	"fmt"
	"io/ioutil"
	"sync"
)

var waitgroup sync.WaitGroup

type Result struct {
	Filename   string
	ImportList ast.ImportList
	Err        error
}

func parseFile(filename string, result *Result) {
	// We expect to be run concurrently behind 'waitgroup'.
	defer waitgroup.Done()

	result.Filename = filename

	code, err := ioutil.ReadFile(filename)
	if err != nil {
		result.Err = err
		return
	}

	// Create a parser with the Context property our ast expects to see. You don't have to
	// that the context be visible to the ast as anything other than an interface{}. It's
	// ok for the UserContext to be opaque to the ast if you're ok with having interface{}
	// types in your ast.
	p := parser.NewParser()
	p.UserContext = &ast.Context{Path: ".", Filename: filename}

	// Create the lexer to tokenize the source for the parser.
	l := lexer.NewLexer(code)

	// ast will be nil if the parse fails.
	var importList interface{}
	importList, result.Err = p.Parse(l)
	if result.Err == nil {
		result.ImportList = importList.(ast.ImportList)
	}
}

func main() {
	// Simulate a number of files to parse in parallel.
	var filenames = []string{
		"sample1.txt", "sample2.txt", "sample3.txt", "sample4.txt", "sample5.txt",
	}

	// Somewhere to store results.
	var results = make([]Result, len(filenames))

	// Parse each file with its own worker.
	for idx, filename := range filenames {
		waitgroup.Add(1)
		// Execute parse concurrently.
		go parseFile(filename, &results[idx])
	}

	// Wait for all workers to call .Done()
	waitgroup.Wait()

	good, bad := 0, 0
	for _, result := range results {
		if result.Err != nil {
			// if we were actually implementing imports, then 'result.Filename' might not
			// tell us which file the error was actually in.
			err := result.Err.(*errors.Error)
			pos := err.ErrorToken.Pos
			fmt.Printf("'%s' errored:\n", result.Filename)
			fmt.Printf("%s:%d:%d: error: %s\n", err.UserContext.(*ast.Context).Filename, pos.Line, pos.Column,
				err.Error())
			bad++
		} else {
			fmt.Printf("'%s' parsed:\n", result.Filename)
			for _, imported := range result.ImportList {
				fmt.Printf("imported from %s:%d:%d: '%s'\n",
					imported.Context.Filename, imported.Token.Pos.Line, imported.Token.Pos.Column, string(imported.Token.Lit))
			}
			good++
		}
	}

	fmt.Printf("passed: %d, failed: %d\n", good, bad)
}
