package items

import (
	"fmt"
	"os"
	"regexp"
	"runtime/debug"
	"strings"
)

var (
	callerRE = regexp.MustCompile("^\\s*(.*)/([^/]+)(\\(.*\\))$")
	sourceRE = regexp.MustCompile("^\\s*(.*)/([^/]+:\\d+)\\s*(.*)$")
)

func dbg(set int, format string, v ...interface{}) {
	if set >= 0 {
		fmt.Fprintf(os.Stderr, "S%d:", set)
	}
	fmt.Fprintf(os.Stderr, format, v...)
}

func trace(set int, format string, v ...interface{}) {
	stack := strings.Split(string(debug.Stack()), "\n")
	fmt.Fprintf(os.Stderr, "-------------------------------------------\n")
	for ix := 5; ix < len(stack)-3; ix += 2 {
		caller := callerRE.ReplaceAllString(stack[ix], "$2")
		source := sourceRE.ReplaceAllString(stack[ix+1], "$2")
		fmt.Fprintf(os.Stderr, "%s in %s\n", caller, source)
	}
	dbg(set, format, v...)
}
