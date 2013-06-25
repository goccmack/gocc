//Copyright 2012 Vastech SA (PTY) LTD
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

//The gen package writes the files and generates the code.
package gen

import (
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/config"
	"code.google.com/p/gocc/lr1"
	"code.google.com/p/gocc/lr1/gen/gocode"
	"code.google.com/p/gocc/token"
	"bytes"
	"fmt"
	"os"
	"path"
	"strings"
)

//Generates and writes the parser, scanner and token packages, given the generated tables and tokenMap.
// func WriteFiles(debug bool, srcDir, pkg, prjName string, tables string, initDecl string, tm *token.TokenMap, genScanner bool) (err error) {
func WriteFiles(cfg config.Config, tables string, initDecl string, tm *token.TokenMap, actTab lr1.ActionTable, gotoTab *lr1.GotoTable, g *ast.Grammar) (err error) {
	errorsImport := `import errs "` + path.Join(cfg.Package(), "errors") + `"` + "\n\n"
	tokenImport := `import "` + path.Join(cfg.Package(), "token") + `"` + "\n\n"

	if err = writeErrors(cfg, tokenImport); err != nil {
		return
	}
	if err = writeParser(cfg, errorsImport, tokenImport); err != nil {
		return
	}
	if cfg.GenScanner() {
		if err = writeScanner(cfg, tokenImport); err != nil {
			return
		}
	}
	if err = writeFileString(path.Join(cfg.TokenDir(), "token.go"), tokenSrc); err != nil {
		return
	}
	if err = writeTables(cfg, initDecl, tables); err != nil {
		return
	}
	if err = writeUTables(cfg, initDecl, actTab, gotoTab, tm, g); err != nil {
		return
	}
	return writeTokenFile(cfg, cfg.ProjectName(), tm)
}

func writeErrors(cfg config.Config, tokenImport string) error {
	fpath := path.Join(cfg.ErrorsDir(), "errors.go")
	return writeFileString(fpath, errorsSrcHeader, tokenImport, errorsSrcBody)
}

func writeParser(cfg config.Config, errorsImport, tokenImport string) error {
	fpath := path.Join(cfg.ParserDir(), "parser.go")
	if err := writeFileString(fpath, parserSrcHdr, errorsImport, tokenImport, parserSrcBody); err != nil {
		return err
	}
	fpath = path.Join(cfg.ParserDir(), "parser_ut.go")
	return writeFileString(fpath, gocode.ParserUTHdrSrc, errorsImport, tokenImport, gocode.ParserUTBodySrc)
}

func writeScanner(cfg config.Config, tokenImport string) error {
	fpath := path.Join(cfg.ScannerDir(), "scanner.go")
	return writeFileString(fpath, scannerSrcHdr, tokenImport, scannerSrcBody)
}

func writeTables(cfg config.Config, initDecl, tablesSrc string) (err error) {
	fpath := path.Join(cfg.ParserDir(), "tables.go")
	return writeFileString(fpath, tableSrcHdr, initDecl, "\n\n", tablesSrc)
}

func writeUTables(cfg config.Config, initDecl string, actTab lr1.ActionTable, gto *lr1.GotoTable, tm *token.TokenMap, g *ast.Grammar) (err error) {
	fpath := path.Join(cfg.ParserDir(), "tables_uncompressed.go")
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "package parser\n\n")
	fmt.Fprintf(buf, "%s\n\n", initDecl)
	gocode.NewActTableGenerator(actTab, cfg).GenUncompressed(buf, tm)
	gocode.NewUGotoTable(gto, cfg).GenUncompressed(buf)
	gocode.NewUProdsTable(g).GenUncompressed(buf)
	return writeFile(fpath, buf.Bytes())
}

func writeTokenFile(cfg config.Config, prjName string, tm *token.TokenMap) error {
	src := "package token\n\n"
	src += "var " + strings.ToUpper(prjName) + "Tokens = NewMapFromStrings([]string{\n"
	for _, s := range tm.Strings() {
		src += "\t" + `"` + s + `"` + ",\n"
	}
	src += "})\n"
	fpath := path.Join(cfg.TokenDir(), "tokens.go")
	return writeFileString(fpath, src)
}

func writeFileString(fpath string, src ...string) (err error) {
	os.Remove(fpath)
	dir, _ := path.Split(fpath)
	if err = os.MkdirAll(dir, 0777); err != nil {
		return
	}
	file, err := os.Create(fpath)
	if err != nil {
		return err
	}
	for _, s := range src {
		if _, err = file.Write([]byte(s)); err != nil {
			file.Close()
			return
		}
	}
	err = file.Close()
	return
}

func writeFile(fpath string, src ...[]byte) (err error) {
	os.Remove(fpath)
	dir, _ := path.Split(fpath)
	if err = os.MkdirAll(dir, 0777); err != nil {
		return
	}
	file, err := os.Create(fpath)
	if err != nil {
		return err
	}
	for _, s := range src {
		if _, err = file.Write(s); err != nil {
			file.Close()
			return
		}
	}
	err = file.Close()
	return
}
