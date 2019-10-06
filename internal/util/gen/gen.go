package gen

import (
	"github.com/maxcalandrelli/gocc/internal/util/gen/golang"
)

func Gen(outDir, subpath string) {
	golang.GenRune(outDir, subpath)
	golang.GenLitConv(outDir, subpath)
}
