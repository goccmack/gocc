//
//  Syntax Directed Translation helper
//

package ast

import (
	"regexp"
	"strings"

	"github.com/maxcalandrelli/gocc/internal/frontend/token"
)

func SDTVal(T *token.Token) string {
	sdt := string(T.Lit)
	rex, err := regexp.Compile("\\$[0-9]+")
	if err != nil {
		panic(err)
	}
	idx := rex.FindAllStringIndex(sdt, -1)
	res := ""
	if len(idx) <= 0 {
		res = sdt
	} else {
		for i, loc := range idx {
			if loc[0] > 0 {
				if i > 0 {
					res += sdt[idx[i-1][1]:loc[0]]
				} else {
					res += sdt[0:loc[0]]
				}
			}
			res += "X["
			res += sdt[loc[0]+1 : loc[1]]
			res += "]"
		}
		if idx[len(idx)-1][1] < len(sdt) {
			res += sdt[idx[len(idx)-1][1]:]
		}
	}
	return strings.TrimSpace(res[2 : len(res)-2])
}
