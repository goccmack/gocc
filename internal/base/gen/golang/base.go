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

package golang

import (
	"bytes"
	"go/format"
	"path"
	"text/template"

	"github.com/maxcalandrelli/gocc/internal/io"
)

type data struct {
	MyName string
}

func GenBase(pkg, outdir string) {
	baseName := path.Base(outdir)
	tokenPath := path.Join(outdir, baseName+".go")
	tmpl, err := template.New(baseName).Parse(TokenMapSrc[1:])
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data{MyName: baseName})
	// Use go/format to indent the idMap literal correctly.
	source, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	io.WriteFile(tokenPath, source)
}

const TokenMapSrc string = `
// Code generated by gocc; DO NOT EDIT.

package {{.MyName}}


`
