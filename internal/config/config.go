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
	"go/build"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	VERSION = "2.1.1001"

	INTERNAL_SYMBOL_EMPTY   = "ε"
	INTERNAL_SYMBOL_ERROR   = "λ"
	INTERNAL_SYMBOL_INVALID = "INVALID"
	INTERNAL_SYMBOL_EOF     = "Ω"
	INTERNAL_SYMBOL_PROD    = "Π"
	INTERNAL_SYMBOL_LIT     = "Λ"
	INTERNAL_SYMBOL_CDTOK   = "μ"
	SYMBOL_EMPTY            = "ε<empty>"
	SYMBOL_ERROR            = "λ<error>"
	SYMBOL_INVALID          = "ά<INVALID>"
	SYMBOL_EOF              = "Ω<EOF>"
	SYMBOL_CDTOK            = "μ"
)

type (
	Config interface {
		Help() bool
		Verbose() bool
		AllowUnreachable() bool
		AutoResolveLRConf() bool
		SourceFile() string
		OutDir() string
		InternalSubdir() string

		NoLexer() bool
		DebugLexer() bool
		DebugParser() bool

		ErrorsDir() string
		ParserDir() string
		ScannerDir() string
		TokenDir() string

		ProjectName() string
		Package() string
		PreProcessor() string

		PrintParams()

		BugOption(string) bugOption
	}

	bugOption  string
	bugOptions map[string]bugOption

	ConfigRecord struct {
		workingDir string

		allowUnreachable  *bool
		autoResolveLRConf *bool
		debugLexer        *bool
		debugParser       *bool
		help              *bool
		noLexer           *bool
		outDir            string
		pkg               string
		srcFile           string
		internal          string
		verbose           *bool
		bug_options       bugOptions
		preprocessor      string
	}
)

const (
	bugopt_fix     = "fix"
	bugopt_ignore  = "ignore"
	bugopt_default = bugopt_fix
)

func (this bugOption) Fix() bool {
	return strings.ToLower(string(this)) != bugopt_ignore
}

func (this bugOption) Ignore() bool {
	return strings.ToLower(string(this)) == bugopt_ignore
}

func (this bugOptions) String() string {
	b := &strings.Builder{}
	for k, _ := range bugDescriptions {
		v := bugDefaultOption
		if _v, _f := this[k]; _f {
			v = _v
		}
		if b.Len() > 0 {
			b.WriteByte(',')
		}
		fmt.Fprintf(b, "%s:%s", k, v)
	}
	return b.String()
}

func (this bugOption) Option() string {
	return string(this)
}

func (this bugOptions) Set(s string) error {
	for _, s1 := range strings.Split(s, ",") {
		s1v := strings.Split(s1, ":")
		if len(s1v) < 2 {
			s1v = append(s1v, bugopt_default)
		}
		s1v[0] = strings.ToLower(s1v[0])
		if _, found := bugDescriptions[s1v[0]]; !found {
			return errors.New("unknown bug name: " + s1v[0])
		}
		this[s1v[0]] = bugOption(strings.Join(s1v[1:], "="))
	}
	return nil
}

var (
	CurrentConfiguration Config
	bugDefaultOption     = bugOption(bugopt_default)
	bugDescriptions      = map[string]string{
		"lexer_dots":    "dots handling in regexps like \"--.--\" needs propagation of fallback states",
		"lexer_regdefs": "incorrect calculation of symbol classes when both a shift and a reduce item exist for the same regdefid",
	}
)

func New() (Config, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	cfg := &ConfigRecord{
		workingDir:  wd,
		bug_options: map[string]bugOption{},
	}

	if err := cfg.getFlags(); err != nil {
		cfg.PrintParams()
		return nil, err
	}

	return cfg, nil
}

func (this *ConfigRecord) Help() bool {
	return *this.help
}

func (this *ConfigRecord) Verbose() bool {
	return *this.verbose
}

func (this *ConfigRecord) AllowUnreachable() bool {
	return *this.allowUnreachable
}

func (this *ConfigRecord) AutoResolveLRConf() bool {
	return *this.autoResolveLRConf
}

func (this *ConfigRecord) NoLexer() bool {
	return *this.noLexer
}

func (this *ConfigRecord) DebugLexer() bool {
	return *this.debugLexer
}

func (this *ConfigRecord) DebugParser() bool {
	return *this.debugParser
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

func (this *ConfigRecord) InternalSubdir() string {
	return this.internal
}

func (this *ConfigRecord) Package() string {
	return this.pkg
}

func (this *ConfigRecord) PreProcessor() string {
	return this.preprocessor
}

func (this *ConfigRecord) ProjectName() string {
	_, file := path.Split(this.srcFile)
	file = file[:len(file)-len(path.Ext(file))]
	return file
}

func (this *ConfigRecord) BugOption(name string) bugOption {
	if _, exists := bugDescriptions[name]; !exists {
		panic("unknown bug name: " + name)
	}
	if option, exists := this.bug_options[name]; exists {
		return option
	}
	return bugDefaultOption
}

func (this *ConfigRecord) PrintParams() {
	fmt.Printf("-a             = %v\n", *this.autoResolveLRConf)
	fmt.Printf("-debug_lexer   = %v\n", *this.debugLexer)
	fmt.Printf("-debug_parser  = %v\n", *this.debugParser)
	fmt.Printf("-h             = %v\n", *this.help)
	fmt.Printf("-no_lexer      = %v\n", *this.noLexer)
	fmt.Printf("-o             = %v\n", this.outDir)
	fmt.Printf("-p             = %v\n", this.pkg)
	fmt.Printf("-u             = %v\n", *this.allowUnreachable)
	fmt.Printf("-v             = %v\n", *this.verbose)
	fmt.Printf("-internal      = %v\n", this.internal)
	fmt.Printf("-bugs          = %v\n", this.bug_options.String())
	fmt.Printf("-preprocessor  = %v\n", this.preprocessor)
}

/*** Utility routines ***/

func bugsHelp(pref string) string {
	b := &strings.Builder{}
	nl := func() {
		b.WriteByte('\n')
		b.WriteString(pref)
	}
	nl()
	fmt.Fprintf(b, "use if you suspect that bug fixing strategies are causing other problems")
	nl()
	fmt.Fprintf(b, "use either --bugs=<bugopt1> --bugs=<bugopt2> or --bugs=<bugopt1>,<bugopt2> form")
	nl()
	fmt.Fprintf(b, "each <bugopt> can be specified as  <bugname> or <bugname>:<bugaction>")
	nl()
	fmt.Fprintf(b, "<bugaction> can be either 'fix' or 'ignore'; some bugs will also allow more options")
	nl()
	fmt.Fprintf(b, "the actually fixable bugs are:")
	nl()
	for name, descr := range bugDescriptions {
		nl()
		fmt.Fprintf(b, "  %-20s %s", name, descr)
	}
	nl()
	nl()
	fmt.Fprintf(b, "example:   gocc -v -bugs=lexer_dots:ignore myfile.bnf")
	nl()
	return b.String()
}

func (this *ConfigRecord) getFlags() error {
	this.autoResolveLRConf = flag.Bool("a", true, "automatically resolve LR(1) conflicts")
	this.debugLexer = flag.Bool("debug_lexer", false, "enable debug logging in lexer")
	this.debugParser = flag.Bool("debug_parser", false, "enable debug logging in parser")
	this.help = flag.Bool("h", false, "help")
	this.noLexer = flag.Bool("no_lexer", false, "do not generate a lexer")
	flag.StringVar(&this.outDir, "o", path.Join(this.workingDir, "@f.grammar", "@f"), "output directory format (@f='name' if input file is 'name.bnf')")
	flag.StringVar(&this.pkg, "p", "", "package, empty defaults to "+defaultPackage(this.outDir))
	flag.StringVar(&this.internal, "internal", "internal", "internal subdir name")
	flag.StringVar(&this.preprocessor, "preprocessor", "none", "preprocessor: 'none','internal', or any command string with placeholders @in and @out")
	flag.Var(this.bug_options, "bugs", "handle bugs in original implementation (default: fix all)"+bugsHelp("  "))
	this.allowUnreachable = flag.Bool("u", false, "allow unreachable productions")
	this.verbose = flag.Bool("v", false, "verbose")
	flag.Parse()

	if *this.noLexer && *this.debugLexer {
		return errors.New("no_lexer and debug_lexer cannot both be set")
	}
	if len(flag.Args()) != 1 && !*this.help {
		return errors.New("Too few arguments")
	}

	this.srcFile = flag.Arg(0)
	this.outDir = getOutDir(this.outDir, this.workingDir, this.srcFile)
	if this.pkg == "" {
		this.pkg = defaultPackage(this.outDir)
	}
	this.pkg = actualize(this.pkg, this.srcFile)
	return nil
}

func actualize(pattern, src string) string {
	_, fname := path.Split(src)
	fname = fname[:len(fname)-len(path.Ext(fname))]
	return regexp.MustCompile("@f").ReplaceAllString(pattern, fname)
}

func getOutDir(outDirSpec, wd, src string) string {
	pattern := func() string {
		if strings.HasPrefix(outDirSpec, wd) {
			return outDirSpec
		}
		if path.IsAbs(outDirSpec) {
			return outDirSpec
		}
		return path.Join(wd, outDirSpec)
	}()
	return actualize(pattern, src)
}

func defaultPackage(wd string) string {
	for _, srcDir := range build.Default.SrcDirs() {
		if strings.HasPrefix(wd, srcDir) {
			pkg := strings.Replace(wd, srcDir, "", -1)
			return sanitizePackage(pkg)
		}
	}
	return sanitizePackage(wd)
}

func sanitizePackage(pkg string) string {
	pkg = filepath.ToSlash(pkg)
	if strings.HasPrefix(pkg, "/") {
		pkg = pkg[1:]
	}
	return pkg
}
