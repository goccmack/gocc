package gen

import (
	"github.com/maxcalandrelli/gocc/internal/util/gen/golang"
)

func Gen(outDir, internal string) {
	golang.GenRune(outDir, internal)
	golang.GenLitConv(outDir, internal)
}
