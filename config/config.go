package config

import(
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

type Config interface{
	Help() bool
	Verbose() bool
	AllowUnreachable() bool
	AutoLRConfResolve() bool
	GenScanner() bool
	Profile() bool
	SourceFile() string
	OutDir() string

	ErrorsDir() string
	ParserDir() string
	ScannerDir() string
	TokenDir() string

	ProjectName() string
	Package() string

	PrintParams()
}

type ConfigRecord struct {
	workingDir string

	allowUnreachable  *bool
	autoLRConfResolve *bool
	profile           *bool
	genScanner        *bool
	help              *bool
	verbose           *bool
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

	if err :=  cfg.getFlags(); err != nil {
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

func (this *ConfigRecord) AutoLRConfResolve() bool {
	return *this.autoLRConfResolve
}

func (this *ConfigRecord) GenScanner() bool {
	return *this.genScanner
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
	fmt.Printf("    resolve LR(1) conflicts       = %t\n", *this.autoLRConfResolve)
	fmt.Printf("    output directory              = %s\n", this.outDir)
	fmt.Printf("    package                       = %s\n", this.pkg)
	fmt.Printf("    generate a scanner            = %t\n", *this.genScanner)
	fmt.Printf("    help                          = %t\n", *this.help)
	fmt.Printf("    allow unreachable productions = %t\n", *this.allowUnreachable)
	fmt.Printf("    resolve LR(1) conflicts       = %t\n", *this.autoLRConfResolve)
	fmt.Printf("    verbose                       = %t\n", *this.verbose)
}


/*** Utility routines ***/

func (this *ConfigRecord) getFlags() error {
	this.allowUnreachable = flag.Bool("u", false, "allow unreachable productions")
	this.autoLRConfResolve = flag.Bool("a", false, "automatically resolve LR(1) conflicts")
	this.genScanner = flag.Bool("s", false, "generate a scanner")
	this.help = flag.Bool("h", false, "help")
	this.verbose = flag.Bool("v", false, "verbose")
	this.profile = flag.Bool("prof", false, "write profile to file")
	flag.StringVar(&this.outDir, "o", this.workingDir, "output dir.")
	flag.StringVar(&this.pkg, "p", defaultPackage(this.outDir), "package")
	flag.Parse()

	if len(flag.Args()) != 1 && !*this.help {
		return errors.New("Too few arguments")
	}

	this.srcFile = flag.Arg(0)

	return nil
}

func defaultPackage(wd string) string {
	srcPath := path.Join(os.Getenv("GOPATH"), "src")
	pkg := strings.Replace(wd, srcPath, "", -1)
	if strings.HasPrefix(pkg, "/") {
		pkg = pkg[1:]
	}
	return pkg
}

