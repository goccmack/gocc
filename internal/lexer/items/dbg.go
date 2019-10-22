package items

import (
	"fmt"
	"os"
	"regexp"
	"runtime/debug"
	"strings"

	"github.com/maxcalandrelli/gocc/internal/util"
)

const (
	debug_wrong_regdef_handling = false
	debug_deeply                = debug_wrong_regdef_handling
)

var (
	callerRE = regexp.MustCompile("^\\s*(.*)/([^/]+)(\\(.*\\))$")
	sourceRE = regexp.MustCompile("^\\s*(.*)/([^/]+:\\d+)\\s*(.*)$")
)

func dbgr(set int, format string, v ...interface{}) {
	if debug_deeply {
		if set >= 0 {
			fmt.Fprintf(os.Stderr, "S%d:", set)
		}
		fmt.Fprintf(os.Stderr, format, v...)
	}
}

func printItems(items ItemList) string {
	b := strings.Builder{}
	for _, i := range items {
		b.WriteString(i.String())
		b.WriteByte('\n')
	}
	return b.String()
}

func dbg(set int, format string, v ...interface{}) {
	if debug_deeply {
		args := make([]interface{}, len(v))
		for _i, _v := range v {
			switch _v.(type) {
			case string:
				_v = util.EscapedString(_v.(string)).Unescape()
			case ItemList:
				_v = printItems(_v.(ItemList))
			case []*Item:
				_v = printItems(ItemList(_v.([]*Item)))
			}
			args[_i] = _v
		}
		dbgr(set, format, args...)
	}
}

func trace(set int, format string, v ...interface{}) {
	if debug_deeply {
		stack := strings.Split(string(debug.Stack()), "\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "-------------------------------------------\n")
		for ix := 5; ix < len(stack)-3; ix += 2 {
			caller := callerRE.ReplaceAllString(stack[ix], "$2")
			source := sourceRE.ReplaceAllString(stack[ix+1], "$2")
			fmt.Fprintf(os.Stderr, "%s in %s\n", caller, source)
		}
		fmt.Fprintf(os.Stderr, "-------------------------------------------\n")
		dbg(set, format, v...)
	}
}
