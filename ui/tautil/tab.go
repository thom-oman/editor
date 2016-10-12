package tautil

func TabRight(ta Texta) {
	_ = alterSelectedText(ta, tabRightLines)
}
func tabRightLines(str string) (string, bool) {
	// assume it's at a line start
	for i := 0; i < len(str); {
		str = str[:i] + string('\t') + str[i:] // insert at start of line
		i = lineEndIndexNextIndex(str, i)
	}
	return str, true
}
func TabLeft(ta Texta) {
	_ = alterSelectedText(ta, tabLeftLines)
}
func tabLeftLines(str string) (string, bool) {
	// assume it's at a line start
	altered := false
	for i := 0; i < len(str); {
		if str[i] == '\t' || str[i] == ' ' {
			// remove
			altered = true
			str = str[:i] + str[i+1:] // +1 is length of '\t' or ' '
		}
		i = lineEndIndexNextIndex(str, i)
	}
	if !altered {
		return "", false
	}
	return str, true
}