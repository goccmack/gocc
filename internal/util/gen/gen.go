package gen

import (
	"github.com/maxcalandrelli/gocc/internal/util/gen/golang"
)

func Gen(outDir string) {
	golang.GenRune(outDir)
	golang.GenLitConv(outDir)
}
