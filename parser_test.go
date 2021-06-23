package parser

import (
	"strings"
	"testing"
)

func Test_Parser(t *testing.T) {
	p := func(str string) *Result {
		if strings.ContainsAny(str[0:1], "abc") {
			return &Result{Token: str[0:1], Count: 1}
		}
		return nil
	}

	arr := Parse("a.bc.a", []Parser{p}).ToArray()

	strCnt := 0
	etcCnt := 0

	for i := 0; i < len(arr); i++ {
		if _, ok := arr[i].Token.(string); ok {
			strCnt += 1
		} else {
			etcCnt += 1
		}
	}
	if strCnt != 4 {
		t.Fail()
	}
	if etcCnt != 2 {
		t.Fail()
	}
}
