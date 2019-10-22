package macro

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

type (
	Macro struct {
		Name          string
		ParametersMap map[string]int
		Text          []string
		EndDelimiter  *regexp.Regexp
		HasVarArgs    bool
	}
	void struct{}
)

var (
	NOTHING            void = void{}
	S_MACRO_ARGS_SEP        = ","
	S_MACRO_ESCAPE_SEP      = fmt.Sprintf("<%s>", S_MACRO_ARGS_SEP)
	S_MACRONOTSEP           = fmt.Sprintf("(?:(?:%s)|[^%s])*", S_MACRO_ESCAPE_SEP, S_MACRO_ARGS_SEP)
	S_MACROPARAMNAME        = "\\b(?:(?:[[:alpha:]](?:[\\w_])*)|(?:VARGS))\\b"
	S_MACROPARAMVALUE       = "\\s*(" + S_MACRONOTSEP + ")\\s*"
	S_MACRONAME             = S_MACROPARAMNAME
	S_MACROARG              = S_MACROPARAMVALUE
	S_MACROARGS             = fmt.Sprintf("(\\s*%s\\s*(?:%s\\s*%s\\s*)*)", S_MACROPARAMNAME, S_MACRO_ARGS_SEP, S_MACROPARAMNAME)
	MACROCONDLINE           = regexp.MustCompile(fmt.Sprintf("^\\s*\\?\\s*(%s)\\s*(==|!=|>|<|>=|<=)\\s*([^\\s]+)\\s+(.+)$", S_MACROPARAMNAME))
	_MACROCONDLINE          = regexp.MustCompile(fmt.Sprintf("^\\s*\\?"))
	MACRONAME               = regexp.MustCompile(fmt.Sprintf("(%s)", S_MACRONAME))
	MACROPARAM              = regexp.MustCompile(fmt.Sprintf("(%s)", S_MACROPARAMNAME))
	MACROARG                = regexp.MustCompile(S_MACROARG)
	MACROARGS               = regexp.MustCompile(S_MACROARGS)
	MACROSTART              = regexp.MustCompile(fmt.Sprintf("^\\s*MACRO\\s+(%s)\\s*(?:\\(%s\\))?\\s*$", S_MACRONAME, S_MACROARGS))
	MACROINVOCATION         = regexp.MustCompile(fmt.Sprintf("\\b(%s)\\s*\\((%s(%s%s)*)\\)\\s*$", S_MACRONAME, S_MACRONOTSEP, S_MACRO_ARGS_SEP, S_MACRONOTSEP))
	MACROVAREXPANSION       = regexp.MustCompile(fmt.Sprintf("{(%s)}", S_MACROPARAMNAME))
	SYNKWD                  = regexp.MustCompile(fmt.Sprintf("\\bKWD\\((%s)((?:,!?%s)*)\\)(\\s|$)", S_MACROPARAMNAME, S_MACROPARAMNAME))
	KWDGRP                  = regexp.MustCompile(fmt.Sprintf(",(!?%s)*", S_MACROPARAMNAME))

	Macros        = map[string]*Macro{}
	Keywords      = map[string]map[string]void{}
	KeywordGroups = map[string]map[string]void{"ALL": {}}
)

func (m *Macro) ArgValue(name string, args []string) (string, error) {
	if index, found := m.ParametersMap[name]; found {
		if index < len(args) {
			return args[index], nil
		} else {
			return "", nil
		}
	}
	if m.HasVarArgs {
		switch name {
		case "VLEN":
			return fmt.Sprintf("%d", len(args)-len(m.ParametersMap)), nil
		case "VARG0":
			if len(args) > len(m.ParametersMap) {
				return args[len(m.ParametersMap)], nil
			}
			return "", nil
		case "VARGS":
			if len(args) > len(m.ParametersMap)+1 {
				return strings.Join(args[len(m.ParametersMap)+1:], ","), nil
			}
			return "", nil
		}
	}
	return "", fmt.Errorf("macro <%s> has no parameter named <%s>", m.Name, name)
}

func (m *Macro) Expand(args []string, out io.Writer) {
	for _, line := range m.Text {
		fmt.Fprintln(out, MACROVAREXPANSION.ReplaceAllStringFunc(line, func(match string) string {
			match = MACROVAREXPANSION.FindStringSubmatch(match)[1]
			//fmt.Printf("args: %q match=<%s> =>%d\n", args, match, m.ParametersMap[match])
			if value, err := m.ArgValue(match, args); err == nil {
				return value
			} else {
				panic(fmt.Errorf("macro <%s> has no parameter named <%s>", m.Name, match))
			}
		}))
	}
}

func processEbnfStream(expanding *Macro, expargs []string, in io.Reader, out io.Writer) {
	currentMacro := (*Macro)(nil)
	lineNumber := 0
	for scanner := bufio.NewScanner(in); scanner.Scan(); {
		lineNumber++
		line := scanner.Text()
		switch {
		case currentMacro != nil:
			if currentMacro.EndDelimiter.MatchString(line) {
				Macros[currentMacro.Name] = currentMacro
				currentMacro = nil
			} else {
				currentMacro.Text = append(currentMacro.Text, line)
				//fmt.Printf("<%s>%03d: %s\n", currentMacro.EndDelimiter.String(), lineNumber, line)
			}
		case currentMacro == nil:
			if cond := MACROCONDLINE.FindStringSubmatch(line); cond != nil {
				if expanding == nil {
					panic(fmt.Errorf("conditional line out of macro"))
				}
				name := cond[1]
				op := cond[2]
				cmp := cond[3]
				rest := cond[4]
				if cmp == "\"\"" || cmp == "''" {
					cmp = ""
				}
				if value, err := expanding.ArgValue(name, expargs); err != nil {
					panic(err)
				} else {
					if func() bool {
						switch op {
						case "==":
							return value == cmp
						case "!=":
							return value != cmp
						case "<":
							return value < cmp
						case ">":
							return value > cmp
						case "<=":
							return value <= cmp
						case ">=":
							return value >= cmp
						case "~=":
							return regexp.MustCompile(cmp).MatchString(value)
						}
						return false
					}() {
						line = rest
					} else {
						line = ""
						continue
					}
				}
			}
			if m := MACROSTART.FindStringSubmatch(line); m != nil {
				macro_name := m[1]
				macro_args := m[2]
				if _, found := Macros[macro_name]; found {
					panic(fmt.Sprintf("macro %s redefined at line %d\n", macro_name, lineNumber))
				}
				currentMacro = &Macro{
					Name:          macro_name,
					ParametersMap: map[string]int{},
					EndDelimiter:  regexp.MustCompile(fmt.Sprintf("^\\s*ENDMACRO\\s*%s\\s*$", macro_name)),
				}
				for _, s := range MACROPARAM.FindAllStringSubmatch(macro_args, -1) {
					name := s[1]
					if currentMacro.HasVarArgs {
						panic(fmt.Errorf("macro <%s>: varargs can only occur as last formal parameter", currentMacro.Name))
					}
					if name == "VARGS" {
						currentMacro.HasVarArgs = true
					} else {
						if _, exists := currentMacro.ParametersMap[name]; exists {
							panic(fmt.Sprintf("macro <%s>: parameter <%s> specified twice\n", currentMacro.Name, name))
						}
						currentMacro.ParametersMap[name] = len(currentMacro.ParametersMap)
					}
				}
				Macros[currentMacro.Name] = currentMacro
				//fmt.Printf("%s parmmap: %q\n", macro_name, currentMacro.ParametersMap)
			} else {
				line = MACROINVOCATION.ReplaceAllStringFunc(line, func(m string) string {
					if vars := MACROINVOCATION.FindStringSubmatch(m); vars != nil {
						macro_name := vars[1]
						macro_args := vars[2]
						if macro, found := Macros[macro_name]; found {
							var builder strings.Builder
							var args []string
							for _, v := range MACROARG.FindAllStringSubmatch(macro_args, -1) {
								v[1] = strings.ReplaceAll(v[1], S_MACRO_ESCAPE_SEP, S_MACRO_ARGS_SEP)
								//fmt.Printf("macro repl: %s vars=%q, v=%q\n", m, vars, v)
								args = append(args, v[1])
							}
							for len(args) > len(macro.ParametersMap) && args[len(args)-1] == "" {
								args = args[:len(args)-1]
							}
							if len(args) == len(macro.ParametersMap) || macro.HasVarArgs && len(args) > len(macro.ParametersMap) {
								macro.Expand(args, &builder)
								processEbnfStream(macro, args, strings.NewReader(builder.String()), out)
								//return builder.String()
							} else {
								panic(fmt.Sprintf("macro <%s> wrong args=<%s> (expected length=%d, actual=%d)\n", macro_name, macro_args, len(macro.ParametersMap), len(args)))
							}
						} else {
							//panic(fmt.Sprintf("unknown macro <%s> in match=<%s>\n", macro_name, m))
							return m
						}
					} else {
						panic(fmt.Sprintf("??? match=<%s>\n", m))
					}
					return ""
				})
				if match := SYNKWD.FindStringSubmatch(line); match != nil {
					kwd := match[1]
					if Keywords[kwd] == nil {
						Keywords[kwd] = map[string]void{"ALL": NOTHING}
						KeywordGroups["ALL"][kwd] = NOTHING
					}
					for _, grp := range KWDGRP.FindAllStringSubmatch(match[2], -1) {
						negate := false
						grpname := grp[1]
						if negate = (grpname[0] == '!'); negate {
							grpname = grpname[1:]
						}
						if KeywordGroups[grpname] == nil {
							KeywordGroups[grpname] = map[string]void{}
						}
						if negate {
							delete(Keywords[kwd], grpname)
							delete(KeywordGroups[grpname], kwd)
						} else {
							Keywords[kwd][grpname] = NOTHING
							KeywordGroups[grpname][kwd] = NOTHING
						}
					}
					line = SYNKWD.ReplaceAllString(line, "KWD_"+kwd)
				}
				if len(line) > 0 {
					fmt.Fprintln(out, line)
				}
			}
		}
	}
}

func internalPreProcess(fname, outname string) error {
	if outf, err := os.Create(outname); err != nil {
		panic(fmt.Sprintf("cannot open <%s>: %s\n", outname, err.Error()))
	} else {
		defer outf.Close()
		if inf, err := os.Open(fname); err != nil {
			panic(fmt.Sprintf("cannot open <%s>: %s\n", fname, err.Error()))
		} else {
			fmt.Fprintf(outf, "//\n// Code generated by gocc preprocessor ; DO NOT EDIT.\n//\n")
			processEbnfStream(nil, nil, inf, outf)
			inf.Close()
			if len(Keywords) > 0 {
				for kwd, _ := range Keywords {
					fmt.Fprintf(outf, "KWD_%s : \"%s\" ;\n", kwd, kwd)
				}
				for grp, km := range KeywordGroups {
					fmt.Fprintf(outf,
						"KGRP_%s:\n%s\n;\n",
						grp,
						strings.Join(
							func() (r []string) {
								for kwd, _ := range km {
									r = append(r, fmt.Sprintf("\t\tKWD_%s\n", kwd))
								}
								return
							}(),
							"\t|\n"),
					)
				}
			}
		}
	}
	return nil
}

func matches(filters []*regexp.Regexp, name string) bool {
	for _, re := range filters {
		if re.MatchString(name) {
			return true
		}
	}
	return false
}

func createPath(outname string) {
	dir, _ := filepath.Split(outname)
	if err := os.MkdirAll(dir, 0764); err != nil {
		panic(fmt.Sprintf("cannot create directory <%s>\n", dir))
	}
}

func PreProcess(preprocessor, fname, outname string) error {
	switch preprocessor {
	case "none":
	case "internal":
		createPath(outname)
		return internalPreProcess(fname, outname)
	default:
		createPath(outname)
		if preprocessCmd := strings.Split(strings.ReplaceAll(strings.ReplaceAll(preprocessor, "@in", fname), "@out", outname), " "); len(preprocessCmd) > 0 {
			cmd := exec.Command(preprocessCmd[0], preprocessCmd[1:]...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			return cmd.Run()
		}
		return fmt.Errorf("cannot run empty preprocessor")
	}
	return nil
}
