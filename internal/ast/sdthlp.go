//
//  Syntax Directed Translation helper
//

package ast

import (
	"regexp"
)

var (
	sdtAstSubst = regexp.MustCompile("(?ms)<<(?:\\s|\\n|\\r)*(.+?)(?:\\s|\\n|\\r)*>>")
)

func SDTVal(sdt string) string {
	return sdtAstSubst.ReplaceAllString(sdt, "$1")
}
