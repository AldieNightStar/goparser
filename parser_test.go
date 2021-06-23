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

func Test_Parser_Ony_Three_Tokens(t *testing.T) {
	p := func(str string) *Result {
		if strings.ContainsAny(str[0:1], "abc") {
			return &Result{Token: str[0:1], Count: 1}
		}
		return nil
	}
	arr := Parse("a.bc.a", []Parser{p}).FewArray(3)

	if len(arr) != 3 {
		t.Fatal("Count is not 3.", len(arr))
	}
}

func Test_Parser_Only_Filtered_Tokens(t *testing.T) {
	p := func(str string) *Result {
		if strings.ContainsAny(str[0:1], "abc") {
			return &Result{Token: str[0:1], Count: 1}
		}
		return nil
	}
	arr := Parse("a.bc.a", []Parser{p}).FilterArray(func(r *Result) bool {
		if _, ok := r.Token.(string); ok {
			return true
		}
		return false
	})

	if len(arr) != 4 {
		t.Fatal("Count is not 4.", len(arr))
	}
}

func Test_Parser_Unti_C_Token(t *testing.T) {
	p := func(str string) *Result {
		if strings.ContainsAny(str[0:1], "abc") {
			return &Result{Token: str[0:1], Count: 1}
		}
		return nil
	}
	arr := Parse("a.bc.a", []Parser{p}).UntilArray(
		func(r *Result) bool {
			if t, ok := r.Token.(string); ok {
				return t == "c"
			}
			return false
		},
	)

	if len(arr) != 3 {
		t.Fatal("Count is not 4.", len(arr))
	}
	if tok, ok := arr[2].Token.(string); ok {
		if tok != "b" {
			t.Fatal("Array last element is not 'b'.", tok)
		}
	} else {
		t.Fatal("Type of last token is not valid!")
	}
}
