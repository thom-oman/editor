package tautil

func StartOfString(ta Texta, sel bool) {
	activateSelection(ta, sel)
	ta.SetCursorIndex(0)
	ta.MakeIndexVisible(0)
	deactivateSelectionCheck(ta)
}