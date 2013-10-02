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

package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

type Config interface {
	Help() bool
	ABNF() bool
	AllowUnreachable() bool
	AutoResolveLRConf() bool
	Knuth() bool
	SourceFile() string
	OutDir() string

	DebugLexer() bool
	DebugParser() bool

	ErrorsDir() string
	ParserDir() string
	ScannerDir() string
	TokenDir() string

	ProjectName() string
	Package() string
	PrintParams()
	Verbose() bool
	Verbose2() bool
}

type ConfigRecord struct {
	workingDir string

	abnf              *bool
	allowUnreachable  *bool
	autoResolveLRConf *bool
	debugLexer        *bool
	debugParser       *bool
	knuth             *bool
	profile           *bool
	genScanner        *bool
	help              *bool
	verbose           *bool
	verbose2          *bool
	srcFile           string
	outDir            string
	pkg               string
}

func New() (Config, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	cfg := &ConfigRecord{
		workingDir: wd,
	}

	if err := cfg.getFlags(); err != nil {
		cfg.PrintParams()
		return nil, err
	}

	return cfg, nil
}

func (this *ConfigRecord) ABNF() bool {
	return *this.abnf
}

func (this *ConfigRecord) Help() bool {
	return *this.help
}

func (this *ConfigRecord) Verbose() bool {
	return *this.verbose
}

func (this *ConfigRecord) Verbose2() bool {
	return *this.verbose2
}

func (this *ConfigRecord) AllowUnreachable() bool {
	return *this.allowUnreachable
}

func (this *ConfigRecord) AutoResolveLRConf() bool {
	return *this.autoResolveLRConf
}

func (this *ConfigRecord) DebugLexer() bool {
	return *this.debugLexer
}

func (this *ConfigRecord) DebugParser() bool {
	return *this.debugParser
}

func (this *ConfigRecord) GenScanner() bool {
	return *this.genScanner
}

func (this *ConfigRecord) Knuth() bool {
	return *this.knuth
}

func (this *ConfigRecord) Profile() bool {
	return *this.profile
}

func (this *ConfigRecord) SourceFile() string {
	return this.srcFile
}

func (this *ConfigRecord) OutDir() string {
	return this.outDir
}

func (this *ConfigRecord) ErrorsDir() string {
	return path.Join(this.outDir, "errors")
}

func (this *ConfigRecord) ParserDir() string {
	return path.Join(this.outDir, "parser")
}

func (this *ConfigRecord) ScannerDir() string {
	return path.Join(this.outDir, "scanner")
}

func (this *ConfigRecord) TokenDir() string {
	return path.Join(this.outDir, "token")
}

func (this *ConfigRecord) Package() string {
	return this.pkg
}

func (this *ConfigRecord) ProjectName() string {
	_, file := path.Split(this.pkg)
	return file
}

func (this *ConfigRecord) PrintParams() {
	fmt.Printf("    generate an ABNF parser       = %t\n", *this.abnf)
	fmt.Printf("    generate debug lexer          = %t\n", *this.debugLexer)
	fmt.Printf("    generate debug parser         = %t\n", *this.debugParser)
	fmt.Printf("    resolve LR(1) conflicts       = %t\n", *this.autoResolveLRConf)
	fmt.Printf("    output directory              = %s\n", this.outDir)
	fmt.Printf("    knuth                         = %t\n", *this.knuth)
	fmt.Printf("    package                       = %s\n", this.pkg)
	fmt.Printf("    generate a scanner            = %t\n", *this.genScanner)
	fmt.Printf("    help                          = %t\n", *this.help)
	fmt.Printf("    allow unreachable productions = %t\n", *this.allowUnreachable)
	fmt.Printf("    resolve LR(1) conflicts       = %t\n", *this.autoResolveLRConf)
	fmt.Printf("    verbose                       = %t\n", *this.verbose)
	fmt.Printf("    more verbose                  = %t\n", *this.verbose2)
}

/*** Utility routines ***/

func (this *ConfigRecord) getFlags() error {
	this.abnf = flag.Bool("ABNF", false, "generate an ABNF parser")
	this.allowUnreachable = flag.Bool("u", false, "allow unreachable productions")
	this.autoResolveLRConf = flag.Bool("a", false, "automatically resolve LR(1) conflicts")
	this.debugLexer = flag.Bool("debug_lexer", false, "generate a debug lexer")
	this.debugParser = flag.Bool("debug_parser", false, "generate a debug parser")
	this.genScanner = flag.Bool("s", false, "generate a scanner")
	this.knuth = flag.Bool("knuth", false, "generate a Knuth LR(1) machine")
	this.help = flag.Bool("h", false, "help")
	this.verbose = flag.Bool("v", false, "verbose")
	this.verbose2 = flag.Bool("vv", false, "more verbose")
	this.profile = flag.Bool("prof", false, "write profile to file")
	flag.StringVar(&this.outDir, "o", this.workingDir, "output dir.")
	flag.StringVar(&this.pkg, "p", defaultPackage(this.outDir), "package")
	flag.Parse()

	this.outDir = getOutDir(this.outDir, this.workingDir)
	if this.outDir != this.workingDir {
		this.pkg = defaultPackage(this.outDir)
	}

	if len(flag.Args()) != 1 && !*this.help {
		return errors.New("Too few arguments")
	}

	this.srcFile = flag.Arg(0)

	return nil
}

func getOutDir(outDirSpec, wd string) string {
	if strings.HasPrefix(outDirSpec, wd) || strings.HasPrefix(outDirSpec, "/") {
		return outDirSpec
	}
	return path.Join(wd, outDirSpec)
}

func defaultPackage(wd string) string {
	srcPath := path.Join(os.Getenv("GOPATH"), "src")
	pkg := strings.Replace(wd, srcPath, "", -1)
	if strings.HasPrefix(pkg, "/") {
		pkg = pkg[1:]
	}
	return pkg
}
