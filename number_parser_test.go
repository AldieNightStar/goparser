package parser

import "testing"

func TestNumberParser(t *testing.T) {
	s := "1.33 d"

	res := NumberParser(s)
	if res == nil {
		t.Fatal("Result is null")
	}
	if res.Token == nil {
		t.Fatal("Token is not present")
	}
	if tok, ok := res.Token.(NumberToken); ok {
		if tok.Value != float64(1.33) {
			t.Fatal("Token is not valid value.", tok)
		}
	}

	s = "1.44.11"
	res = NumberParser(s)
	if res == nil {
		t.Fatal("Result 2 is null")
	}
	if res.Token == nil {
		t.Fatal("Token 2 is not present")
	}
	if tok, ok := res.Token.(NumberToken); ok {
		if tok.Value != float64(1.44) {
			t.Fatal("Token 2 is not valid value.", tok)
		}
	}
}
