package parser

import "testing"

func TestFindLineAt(t *testing.T) {
	s := "line 1.\nline 2.\nline 3.\nline 4.\nline 5."
	linesInfo := Util_GetLines(s)

	lineAssert := func(linesInfo []*LineInfo, cnt, lineNum int, shouldBeNull bool) {
		inf := Util_FindLineInfoAt(linesInfo, cnt)
		if shouldBeNull {
			if inf != nil {
				t.Fatal("Not a null. But should be null.", inf.Value)
			} else {
				return
			}
		}
		if inf == nil && !shouldBeNull {
			t.Fatal("String line info is null")
		}
		if inf.Number != lineNum {
			t.Fatalf("Line number [%d] is not right. Should be [%d] for line: %s", lineNum, inf.Number, inf.Value)
		}
	}

	lineAssert(linesInfo, 1, 0, false)
	lineAssert(linesInfo, 9, 1, false)
	lineAssert(linesInfo, 35, 4, false)
	lineAssert(linesInfo, 100, 0, true)
}
