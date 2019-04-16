package parseutil

import (
	"testing"

	"github.com/jmigpin/editor/util/iout/iorw"
)

func TestParseResource1(t *testing.T) {
	s := "AAA /a/b/c.txt AAA"
	testParseResourcePath(t, s, 10, "/a/b/c.txt")
}

func TestParseResource2(t *testing.T) {
	s := "AAA /a/b/c%20c.txt AAA"
	testParseResourcePath(t, s, 9, "/a/b/c%20c.txt")
}

func TestParseResource3(t *testing.T) {
	s := "AAA /a/b/c\\ c.txt AAA"
	testParseResourcePath(t, s, 9, "/a/b/c c.txt")
}

func TestParseResource4(t *testing.T) {
	s := "AAA /a/b/c\\\\ c.txt AAA"
	testParseResourcePath(t, s, 9, "/a/b/c\\")
}

func TestParseResource5(t *testing.T) {
	s := "AAA /a/b/c\\ c.txt AAA"
	testParseResourcePath(t, s, 11, "/a/b/c c.txt") // index in mid escape
}

func TestParseResource5_2(t *testing.T) {
	s := "AAA /a/b/c\\ c.txt AAA"
	testParseResourcePath(t, s, 12, "/a/b/c c.txt")
}

func TestParseResource5_3(t *testing.T) {
	s := "AAA /a/b/c\\ c.txt AAA"
	testParseResourcePath(t, s, 13, "/a/b/c c.txt")
}

func TestParseResource5_4(t *testing.T) {
	s := "/a/b/c\\ c.txt AAA"
	testParseResourcePath(t, s, 0, "/a/b/c c.txt")
}

func TestParseResource5_5(t *testing.T) {
	s := " a\\ b\\ c.txt"
	testParseResourcePath(t, s, 3, "a b c.txt")
}

func TestParseResource6(t *testing.T) {
	s := "AAA /a/b/c.txt\\:a:1:2# AAA"
	testParseResourcePath(t, s, 11, "/a/b/c.txt:a")
	testParseResourceLineCol(t, s, 11, 1, 2)
}

func TestParseResource7(t *testing.T) {
	s := "AAA /a/b/c.txt\\:a:1:#AAA"
	testParseResourcePath(t, s, 11, "/a/b/c.txt:a")
	testParseResourceLineCol(t, s, 11, 1, 0)
}

func TestParseResource8(t *testing.T) {
	s := "/a/b/c:1:2"
	testParseResourcePath(t, s, 0, "/a/b/c")
	testParseResourceLineCol(t, s, 0, 1, 2)
}

func TestParseResource9(t *testing.T) {
	s := "/a/b\\ b/c"
	testParseResourcePath(t, s, 0, "/a/b b/c")
}

func TestParseResource10(t *testing.T) {
	s := "/a/b\\"
	testParseResourcePath(t, s, 0, "/a/b")
}

func TestParseResource11(t *testing.T) {
	s := ": /a/b/c"
	testParseResourcePath(t, s, len(s), "/a/b/c")
}

func TestParseResource12(t *testing.T) {
	s := "//a/b/////c"
	testParseResourcePath(t, s, len(s), "/a/b/c")
}

func TestParseResource13(t *testing.T) {
	s := "(/a/b/c.txt)"
	testParseResourcePath(t, s, 5, "/a/b/c.txt")
}

func TestParseResource14(t *testing.T) {
	s := "[/a/b/c.txt]"
	testParseResourcePath(t, s, 5, "/a/b/c.txt")
}

func TestParseResource15(t *testing.T) {
	s := "</a/b/c.txt>"
	testParseResourcePath(t, s, 5, "/a/b/c.txt")
}

func TestParseResource16(t *testing.T) {
	s := ""
	_, err := ParseResourceStr(s, 0)
	if err == nil {
		t.Fatal("able to parse empty string")
	}
}

func TestParseResource17(t *testing.T) {
	s := "./a/b/c.txt :20"
	testParseResourcePath(t, s, 3, "./a/b/c.txt")
	testParseResourceLineCol(t, s, 0, 0, 0)
}

func TestAddEscapes(t *testing.T) {
	s := "a \\b"
	s2 := AddEscapes(s, '\\', " \\")
	if s2 != "a\\ \\\\b" {
		t.Fatal()
	}
	s3 := RemoveEscapes(s2, '\\')
	if s3 != s {
		t.Fatal()
	}
}

//----------

func TestExpand1(t *testing.T) {
	s := ": /a/b/c"
	rw := iorw.NewBytesReadWriter([]byte(s))
	l, _ := ExpandResourceIndexes(rw, rw.Max(), '\\')
	if l != 2 {
		t.Fatalf("%v", l)
	}
}

//----------

func testParseResourcePath(t *testing.T, str string, index int, estr string) {
	t.Helper()
	u, err := ParseResourceStr(str, index)
	if err != nil {
		t.Fatal(err)
	}
	if u.Path != estr {
		t.Fatalf("%#v", u)
	}
}

func testParseResourceLineCol(t *testing.T, str string, index int, eline, ecol int) {
	t.Helper()
	u, err := ParseResourceStr(str, index)
	if err != nil {
		t.Fatal(err)
	}
	if u.Line != eline || u.Column != ecol {
		t.Fatalf("%v\n%#v", str, u)
	}
}

//----------
