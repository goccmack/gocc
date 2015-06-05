package gen

import (
	"github.com/goccmack/gocc/util/gen/golang"
)

func Gen(outDir string) {
	golang.GenRune(outDir)
	golang.GenLitConv(outDir)
}
