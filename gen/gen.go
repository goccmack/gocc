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
	"code.google.com/p/gocc/token"
	"os"
	"path"
	"strings"
)

//Generates and writes the parser, scanner and token packages, given the generated tables and tokenMap.
// func WriteFiles(debug bool, srcDir, pkg, prjName string, tables string, initDecl string, tm *token.TokenMap, genScanner bool) (err error) {
func WriteFiles(srcDir, pkg, prjName string, tables string, initDecl string, tm *token.TokenMap, genScanner bool) (err error) {
	errorsDir := path.Join(srcDir, "errors")
	errorsImport := `import errs "` + path.Join(pkg, "errors") + `"` + "\n\n"
	parserDir := path.Join(srcDir, "parser")
	scannerDir := path.Join(srcDir, "scanner")
	tokenDir := path.Join(srcDir, "token")
	tokenImport := `import "` + path.Join(pkg, "token") + `"` + "\n\n"

	if err = writeErrors(errorsDir, tokenImport); err != nil {
		return
	}
	if err = writeParser(parserDir, errorsImport, tokenImport); err != nil {
		return
	}
	if genScanner {
		if err = writeScanner(scannerDir, tokenImport); err != nil {
			return
		}
	}
	if err = writeFile(path.Join(tokenDir, "token.go"), tokenSrc); err != nil {
		return
	}
	if err = writeTables(parserDir, initDecl, tables); err != nil {
		return
	}
	return writeTokenFile(tokenDir, prjName, tm)
}

func writeErrors(errorsDir, tokenImport string) error {
	fpath := path.Join(errorsDir, "errors.go")
	return writeFile(fpath, errorsSrcHeader, tokenImport, errorsSrcBody)
}

func writeParser(parserDir, errorsImport, tokenImport string) error {
	fpath := path.Join(parserDir, "parser.go")
	return writeFile(fpath, parserSrcHdr, errorsImport, tokenImport, parserSrcBody)
}

func writeScanner(scannerDir, tokenImport string) error {
	fpath := path.Join(scannerDir, "scanner.go")
	return writeFile(fpath, scannerSrcHdr, tokenImport, scannerSrcBody)
}

func writeTables(parserDir, initDecl, tablesSrc string) (err error) {
	fpath := path.Join(parserDir, "tables.go")
	return writeFile(fpath, tableSrcHdr, initDecl, "\n\n", tablesSrc)
}

func writeTokenFile(tokenDir, prjName string, tm *token.TokenMap) error {
	src := "package token\n\n"
	src += "var " + strings.ToUpper(prjName) + "Tokens = NewMapFromStrings([]string{\n"
	for _, s := range tm.Strings() {
		src += "\t" + `"` + s + `"` + ",\n"
	}
	src += "})\n"
	fpath := path.Join(tokenDir, "tokens.go")
	return writeFile(fpath, src)
}

func writeFile(fpath string, src ...string) (err error) {
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
