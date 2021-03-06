package toolbarparser

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jmigpin/editor/core/parseutil"
	"github.com/jmigpin/editor/util/iout/iorw"
	"github.com/jmigpin/editor/util/osutil"
	"github.com/jmigpin/editor/util/scanutil"
)

//----------

func ParseVars(data *Data) VarMap {
	m := VarMap{}
	for _, part := range data.Parts {
		if len(part.Args) != 1 { // must have 1 arg
			continue
		}
		str := part.Args[0].Str() // parse first arg only
		v, err := ParseVar(str)
		if err != nil {
			continue
		}
		m[v.Name] = v.Value
	}
	return m
}

//----------

type Var struct {
	Name, Value string
}

func ParseVar(str string) (*Var, error) {
	rd := iorw.NewStringReader(str)
	sc := scanutil.NewScanner(rd)
	ru := sc.PeekRune()
	switch ru {
	case '~':
		return parseTildeVar(sc)
	case '$':
		return parseDollarVar(sc)
	}
	return nil, fmt.Errorf("unexpected rune: %v", ru)
}

//----------

func parseTildeVar(sc *scanutil.Scanner) (*Var, error) {
	// name
	if !sc.Match.Sequence("~") {
		return nil, sc.Errorf("name")
	}
	if !sc.Match.Int() {
		return nil, sc.Errorf("name")
	}
	name := sc.Value()
	sc.Advance()
	// assign (must have)
	if !sc.Match.Any("=") {
		return nil, sc.Errorf("assign")
	}
	sc.Advance()
	// value (must have)
	v, err := parseVarValue(sc, false)
	if err != nil {
		return nil, err
	}
	// end
	_ = sc.Match.Spaces()
	if !sc.Match.End() {
		return nil, sc.Errorf("not at end")
	}

	w := &Var{Name: name, Value: v}
	return w, nil
}

//----------

func parseDollarVar(sc *scanutil.Scanner) (*Var, error) {
	// name
	if !sc.Match.Sequence("$") {
		return nil, sc.Errorf("name")
	}
	if !sc.Match.Id() {
		return nil, sc.Errorf("name")
	}
	name := sc.Value()
	sc.Advance()

	w := &Var{Name: name}

	// assign (optional)
	if !sc.Match.Any("=") {
		return w, nil
	}
	sc.Advance()
	// value (optional)
	value, err := parseVarValue(sc, true)
	if err != nil {
		return nil, err
	}
	w.Value = value
	// end
	_ = sc.Match.Spaces()
	if !sc.Match.End() {
		return nil, sc.Errorf("not at end")
	}

	return w, nil
}

//----------

func parseVarValue(sc *scanutil.Scanner, allowEmpty bool) (string, error) {
	if sc.Match.Quoted("\"'", osutil.EscapeRune, true, 1000) {
		v := sc.Value()
		sc.Advance()
		u, err := strconv.Unquote(v)
		if err != nil {
			return "", sc.Errorf("unquote: %v", err)
		}
		return u, nil
	} else {
		if !sc.Match.ExceptUnescapedSpaces(osutil.EscapeRune) {
			if !allowEmpty {
				return "", sc.Errorf("value")
			}
		}
		v := sc.Value()
		sc.Advance()
		return v, nil
	}
}

//----------

type VarMap map[string]string // name -> value

func EncodeVars(filename string, m VarMap) string {
	return parseutil.EscapeFilename(encodeVars(filename, m))
}
func encodeVars(f string, m VarMap) string {
	best := ""
	for k, v := range m {
		v2 := DecodeVars(v, m)

		// exact match
		if f == v2 {
			return k
		}

		// (var + separator) prefix match (best is shortest len)
		v3 := v2 + string(filepath.Separator)
		if strings.HasPrefix(f, v3) {
			s := filepath.Join(k, f[len(v2):])
			if best == "" || len(s) < len(best) {
				best = s
			}
		}
	}
	if best != "" {
		return best
	}
	return f
}

//----------

func DecodeVars(f string, m VarMap) string {
	return parseutil.RemoveEscapes(decodeVars(f, m), osutil.EscapeRune)
}
func decodeVars(f string, m VarMap) string {
	f = filepath.Clean(f)

	// split on first separator
	i := strings.IndexFunc(f, func(ru rune) bool {
		return ru == filepath.Separator
	})
	s0, s1 := f, ""
	if i > 0 {
		s0, s1 = f[:i], f[i:]
	}

	v, ok := m[s0]
	if ok {
		v2 := DecodeVars(v, m)
		return filepath.Join(append([]string{v2}, s1)...)
	}

	return f
}
