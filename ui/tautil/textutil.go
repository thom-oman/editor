package tautil

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func activateSelection(ta Texta, active bool) {
	if active {
		if !ta.SelectionOn() {
			ta.SetSelectionOn(true)
			ta.SetSelectionIndex(ta.CursorIndex())
		}
	} else {
		ta.SetSelectionOn(false)
	}
}
func deactivateSelectionCheck(ta Texta) {
	if ta.SelectionOn() {
		if ta.CursorIndex() == ta.SelectionIndex() {
			ta.SetSelectionOn(false)
		}
	}
}

func isNotSpace(ru rune) bool {
	return !unicode.IsSpace(ru)
}

func NextRuneIndex(str string, index int) (rune, int, bool) {
	ru, size := utf8.DecodeRuneInString(str[index:])
	if ru == utf8.RuneError {
		if size == 0 { // empty string
			return 0, 0, false
		}
		// size==1// invalid encoding, continue with 1
		ru = rune(str[index+size])
	}
	return ru, index + size, true
}
func PreviousRuneIndex(str string, index int) (rune, int, bool) {
	ru, size := utf8.DecodeLastRuneInString(str[:index])
	if ru == utf8.RuneError {
		if size == 0 { // empty string
			return 0, 0, false
		}
		// size==1 // invalid encoding, continue with 1
		ru = rune(str[index-size])
	}
	return ru, index - size, true
}

func selectionStringIndexes(ta Texta) (int, int, bool) {
	if !ta.SelectionOn() {
		panic("!")
	}
	a := ta.SelectionIndex()
	b := ta.CursorIndex()
	if a > b {
		a, b = b, a
	}
	return a, b, a != b
}

func linesStringIndexes(ta Texta) (int, int, bool) {
	var a, b int
	if ta.SelectionOn() {
		var ok bool
		a, b, ok = selectionStringIndexes(ta)
		if !ok {
			return 0, 0, false
		}
	} else {
		a = ta.CursorIndex()
		b = a
	}
	a = lineStartIndex(ta.Text(), a)
	b = lineEndIndexNextIndex(ta.Text(), b)
	return a, b, a != b
}

func lineStartIndex(str string, index int) int {
	i := strings.LastIndex(str[:index], "\n")
	if i < 0 {
		i = 0
	} else {
		i += 1 // rune length of '\n'
	}
	return i
}
func lineEndIndexNextIndex(str string, index int) int {
	i := strings.Index(str[index:], "\n")
	if i < 0 {
		return len(str)
	}
	return index + i + 1 // 1 is "\n" size
}

// used in: comment/uncomment, tabright/tableft
func alterSelectedText(ta Texta, fn func(string) (string, bool)) bool {
	a, b, ok := linesStringIndexes(ta)
	if !ok {
		return false
	}

	t, ok := fn(ta.Text()[a:b])
	if !ok {
		return false
	}

	c := len(t)
	// previous rune so it doesn't include last \n
	if t[len(t)-1] == '\n' {
		_, c2, ok := PreviousRuneIndex(t, len(t))
		if !ok {
			return false
		}
		c = c2
		if c == 0 {
			return false // a==b
		}
	}
	// replace text
	ta.SetText(ta.Text()[:a] + t + ta.Text()[b:])

	// set cursor after the replacement
	ta.SetCursorIndex(a + c)

	// select the replacement
	ta.SetSelectionOn(true)
	ta.SetSelectionIndex(a)

	return true
}