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

package gen

import (
	"code.google.com/p/gocc/token"
	"os"
	"path"
	"strings"
)

func WriteFiles(debug bool, srcDir, pkgRoot, prjName string, tables string, tm *token.TokenMap, genScanner bool) (err error) {
	pkgDir := path.Join(srcDir, pkgRoot)
	parserDir := path.Join(pkgDir, prjName, "parser")
	scannerDir := path.Join(pkgDir, prjName, "scanner")
	tokenDir := path.Join(pkgDir, prjName, "token")
	tokenImport := `import "` + path.Join(pkgRoot, prjName, "token") + `"` + "\n\n"

	if err = writeParser(parserDir, tokenImport, debug); err != nil {
		return
	}
	if genScanner {
		if err = writeScanner(scannerDir, tokenImport); err != nil {
			return
		}
	}
	if err = writeFile(path.Join(pkgDir, prjName, "token", "token.go"), tokenSrc); err != nil {
		return
	}
	if err = writeTables(parserDir, pkgRoot, prjName, tables); err != nil {
		return
	}
	return writeTokenFile(tokenDir, prjName, tm)
}

func writeParser(parserDir, tokenImport string, debug bool) error {
	fpath := path.Join(parserDir, "parser.go")
	var body string
	if debug {
		body = debugParserSrcBody
	} else {
		body = parserSrcBody
	}
	return writeFile(fpath, parserSrcHdr, tokenImport, body)
}

func writeScanner(scannerDir, tokenImport string) error {
	fpath := path.Join(scannerDir, "scanner.go")
	return writeFile(fpath, scannerSrcHdr, tokenImport, scannerSrcBody)
}

func writeTables(parserDir, pkg, prjName, tablesSrc string) (err error) {
	fpath := path.Join(parserDir, "tables.go")
	astImport := `import "` + path.Join(pkg, prjName, "ast") + "\"\n\n"
	return writeFile(fpath, tableSrcHdr, astImport, tablesSrc)
}

func writeTokenFile(scannerDir, prjName string, tm *token.TokenMap) error {
	src := "package token\n\n"
	src += "var " + strings.ToUpper(prjName) + "Tokens = NewMapFromStrings([]string{\n"
	for _, s := range tm.Strings() {
		src += "\t" + `"` + s + `"` + ",\n"
	}
	src += "})\n"
	fpath := path.Join(scannerDir, prjName+"tokens.go")
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
