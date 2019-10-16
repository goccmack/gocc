//
//go:generate go run stock/main.go -a -v -o internal/frontend/reparsed spec/gocc2.ebnf
//

//Copyright 2013 Vastech SA (PTY) LTD
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

//Gocc is LR1 parser generator for go written in go. The generator uses a BNF with very easy to use SDT rules.
//Please see https://github.com/goccmack/gocc/ for more documentation.

/*
   Modified by: Massimiliano Calandrelli <max@maxcalandrelli.it>

  Changes summary:

    2019-10-14
    Many changes, many new features. Streamed parsing, longest prefix parsing, nondeterministic subparsing.

    - changed all go files using import from
        https://github.com/goccmack/gocc
      to import from
        https://github.com/maxcalandrelli/gocc

    - fixed a state machine generation bug (hopefully), that prevented a RE like '<' '<' . { . } '>' '>'
      to recognize a string like "<< a > b >>" (this is done in ItemSets.propagateDots and in the generated
      lexer/transitiontable.go)

    - eliminated ambiguity between literals and labels of lexical or syntaxical productions; "error" can now
      be specified as a literal, and a synonym is "λ", while a synonim for "empty" is "ε"

    - used original gocc (https://github.com/goccmack/gocc) to reproduce the initial (handwritten?) parser

    - changed defaults for output directory

    - changed output directory structure

    - added the ability to specify negative char and char ranges to lexer, as in:

        all_but_star : . | ~'*' ;
        hexdigit : '0'-'F'  | ~(':'-'@') ;

    - added a user defined context to parsing for simpler cases, where an AST is not required. the context is
      available as an implicit variable with name "Context" in SDT snippets. a variable named "Stream" contains
      the underlying lexer stream

    - added functions shorthand substitutions for SDT placeholders $<N>, with the format

                            $<N>:<ops>

      the operations in <ops> are applied to the token/ast; ops is a string where every character stays for an unary
      function of a string, identified by one of the following:
          - s:    converts $<N> to string directly if it is a token in the target grammar; if it is some other object
                  (i.e. an AST branch returned by a reduce operation), it is converted to a string with a .(string) conversion
                  and, in case of failure, with a "%q" format string
          - q:    if the resulting string is enclosed in single o double quotes, it is unwrapped an the quotes are
                  discarded
          - e:    unescape the string, replacing control sequences with their corresponding represented values
          - U:    uppercase conversion
          - l:    lowercase conversion
        for example, if the literal in the token in a terminal symbol in third position in a syntactical production is
        <"'value of PI (\u03c0): \"3.14159\"'"> (angular braces used in this text to quote the values), like in:

            ProdX: HeaderPart OptionalPart quoted_string Trailerpart << ast.PrepareWhatever(...) >> ;

        then we will get the following results for quoted_string ($2):

            $2       *token.Token object whose Lit field has the string value <"'value of PI (\u03c0): \"3.14159\"'">
            $2s      <"'value of PI (\u03c0): \"3.14159\"'">
            $2e      <"'value of PI (π): "3.14159"'">
            $2eq     <'value of PI (π): "3.14159"'>
            $2eqq    <value of PI (π): "3.14159">
            $2eqU    <VALUE OF PI (Π): "3.14159">
            $2eqUq   <VALUE OF PI (Π): "3.14159">

    - added the ability to parse only the longest possible prefix of data, returning the consumed bytes

        subTree, err, parsed := myParser.ParseLongestPrefix(myScanner)

    - added support for a simple form of context-sensitive (non deterministic) parsing by invoking a different
      lexer/parser pair while scanning

      - compact form, using a gocc generated parser for longest prefix sub-parsing:

          NumericValue:
              integer
            |
          		@ "github.com/maxcalandrelli/gocc/example/calc/calc.grammar/calc"
          			<<
          				$1, nil
          			>>

        in compact form, an import alias can be specified (without changes in stack usage)to avoid naming conflicts;
        the following example is totally equivalent to the former:

          NumericValue:
              integer
            |
          		@ calculator "github.com/maxcalandrelli/gocc/example/calc/calc.grammar/calc"
          			<<
          				$1, nil
          			>>

      - full form, like in the following example, where also the use of "Stream" is shown:

          SlashesOrKeyword:
          		"no-slashes"
                << "no slashes", nil >>
          	|
          		somethingelse
          		@@
        				func () (interface {}, error, []byte) {
        					slashes := []byte{}
        					for r, _, _ := Stream.ReadRune(); r == '/' || r == '\\' || r == '\u2215' || r == '\u29f5'; r, _, _ = Stream.ReadRune() {
        						slashes = append(slashes, string(r)...)
        					}
        					Stream.UnreadRune()
        					return len(string(slashes)), nil, slashes
        				}()
          		@@
          		OtherProduction
          			<<
          				append(append([]interface{}{},$1),$2.([]interface{})...), nil
          			>>

        either form of non deterministic parsing pushes two values on the stack:
          - the pseudo-token built by taking as literal the longest prefix that the sub-parser rcognized
          - the corresponding AST returned by the sun-parser



*/

package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/maxcalandrelli/gocc/internal/ast"
	genBase "github.com/maxcalandrelli/gocc/internal/base/gen"
	"github.com/maxcalandrelli/gocc/internal/config"
	altfe "github.com/maxcalandrelli/gocc/internal/frontend/reparsed"
	"github.com/maxcalandrelli/gocc/internal/io"
	genIo "github.com/maxcalandrelli/gocc/internal/io/gen"
	genLexer "github.com/maxcalandrelli/gocc/internal/lexer/gen/golang"
	lexItems "github.com/maxcalandrelli/gocc/internal/lexer/items"
	"github.com/maxcalandrelli/gocc/internal/parser/first"
	genParser "github.com/maxcalandrelli/gocc/internal/parser/gen"
	lr1Action "github.com/maxcalandrelli/gocc/internal/parser/lr1/action"
	lr1Items "github.com/maxcalandrelli/gocc/internal/parser/lr1/items"
	"github.com/maxcalandrelli/gocc/internal/parser/symbols"
	outToken "github.com/maxcalandrelli/gocc/internal/token"
	genToken "github.com/maxcalandrelli/gocc/internal/token/gen"
	genUtil "github.com/maxcalandrelli/gocc/internal/util/gen"
)

func main() {
	flag.Usage = usage
	cfg, err := config.New()
	if err != nil {
		fmt.Printf("Error reading configuration: %s\n", err)
		flag.Usage()
	}

	if cfg.Verbose() {
		fmt.Fprintf(os.Stderr, "gocc version %s\n", config.VERSION)
		cfg.PrintParams()
	}

	if cfg.Help() {
		fmt.Fprintf(os.Stderr, "gocc version %s\n", config.VERSION)
		flag.Usage()
	}

	config.CurrentConfiguration = cfg

	var (
		grammar interface{}
	)
	ast.StringGetter = func(v interface{}) string { return string(v.(*altfe.Token).Lit) }
	grammar, err = altfe.ParseFile(cfg.SourceFile())
	if err != nil {
		fmt.Printf("Parse error: %s\n", err)
		os.Exit(1)
	}

	outdir_base := cfg.OutDir()
	outdir_log := path.Join(outdir_base, "log")
	outdir_iface := path.Join("iface")
	g := grammar.(*ast.Grammar)

	gSymbols := symbols.NewSymbols(g)
	if cfg.Verbose() {
		writeTerminals(gSymbols, cfg, outdir_log)
	}

	var tokenMap *outToken.TokenMap

	gSymbols.Add(g.LexPart.TokenIds()...)
	g.LexPart.UpdateStringLitTokens(gSymbols.ListStringLitSymbols())
	lexSets := lexItems.GetItemSets(g.LexPart)
	if cfg.Verbose() {
		io.WriteFileString(path.Join(outdir_log, "lexer_sets.txt"), lexSets.String())
	}
	tokenMap = outToken.NewTokenMap(gSymbols.ListTerminals())
	if !cfg.NoLexer() {
		genLexer.Gen(cfg.Package(), outdir_base, g.LexPart.Header.SDTLit, lexSets, tokenMap, cfg, cfg.InternalSubdir(), outdir_iface)
	}
	hasSyntax := (g.SyntaxPart != nil)
	if hasSyntax {
		firstSets := first.GetFirstSets(g, gSymbols)
		if cfg.Verbose() {
			io.WriteFileString(path.Join(outdir_log, "first.txt"), firstSets.String())
		}

		lr1Sets := lr1Items.GetItemSets(g, gSymbols, firstSets)
		if cfg.Verbose() {
			io.WriteFileString(path.Join(outdir_log, "LR1_sets.txt"), lr1Sets.String())
		}

		conflicts := genParser.Gen(cfg.Package(), outdir_base, g.SyntaxPart.Header.SDTLit, g.SyntaxPart.ProdList, gSymbols, lr1Sets, tokenMap, cfg, cfg.InternalSubdir(), outdir_iface)
		handleConflicts(conflicts, lr1Sets.Size(), cfg, g.SyntaxPart.ProdList, outdir_log)
	}

	genToken.Gen(cfg.Package(), outdir_base, tokenMap, cfg.InternalSubdir(), cfg)
	genUtil.Gen(outdir_base, cfg.InternalSubdir())
	genBase.Gen(cfg.Package(), outdir_base, cfg.InternalSubdir(), outdir_iface, cfg, hasSyntax)
	genIo.Gen(cfg.Package(), outdir_base, cfg.InternalSubdir())
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: gocc flags bnf_file\n\n")
	fmt.Fprintf(os.Stderr, "  bnf_file: contains the BNF grammar\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func handleConflicts(conflicts map[int]lr1Items.RowConflicts, numSets int, cfg config.Config, prods ast.SyntaxProdList, outdir string) {
	if len(conflicts) <= 0 {
		return
	}
	fmt.Printf("%d LR-1 conflicts \n", len(conflicts))
	if cfg.Verbose() || !cfg.AutoResolveLRConf() {
		io.WriteFileString(path.Join(outdir, "LR1_conflicts.txt"), conflictString(conflicts, numSets, prods))
	}
	if !cfg.AutoResolveLRConf() {
		os.Exit(1)
	}
}

func conflictString(conflicts map[int]lr1Items.RowConflicts, numSets int, prods ast.SyntaxProdList) string {
	w := new(strings.Builder)
	fmt.Fprintf(w, "%d LR-1 conflicts: \n", len(conflicts))
	for i := 0; i < numSets; i++ {
		if cnf, exist := conflicts[i]; exist {
			fmt.Fprintf(w, "\tS%d\n", i)
			for sym, conflicts := range cnf {
				fmt.Fprintf(w, "\t\tsymbol: %s\n", sym)
				for _, cflct := range conflicts {
					switch c := cflct.(type) {
					case lr1Action.Reduce:
						fmt.Fprintf(w, "\t\t\tReduce(%d:%s)\n", c, prods[c])
					case lr1Action.Shift:
						fmt.Fprintf(w, "\t\t\t%s\n", cflct)
					default:
						panic(fmt.Sprintf("unexpected type of action: %s", cflct))
					}
				}
			}
		}
	}
	return w.String()
}

func writeTerminals(gSymbols *symbols.Symbols, cfg config.Config, outdir string) {
	buf := new(bytes.Buffer)
	for _, t := range gSymbols.ListTerminals() {
		fmt.Fprintf(buf, "%s\n", t)
	}
	io.WriteFile(path.Join(outdir, "terminals.txt"), buf.Bytes())
}
