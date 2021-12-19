package parser

import "testing"

func Test_FuncCallParser_positive(t *testing.T) {
	f := "alan(a, b(1, 2), c, 123, \"321\") xxx"
	res := FuncCallParser(f)
	if res == nil {
		t.Fatal("Func call cannot be nil")
	}
	if res.Count != 31 {
		t.Fatal("Count is not 31")
	}
	tok, ok := res.Token.(*FuncCallToken)
	if !ok {
		t.Fatal("FuncCall token isn't FuncCallToken type")
	}
	if tok.FuncName != "alan" {
		t.Fatal("Func call name isn't ok")
	}
	if len(tok.Args) != 5 {
		t.Fatal("Func call args isn't 5")
	}
	if last, lastOk := tok.Args[4].(*StringToken); lastOk {
		if last.Value != "321" {
			t.Fatal("Last token is not valid")
		}
	} else {
		t.Fatal("Last token should be string")
	}
}

func Test_FuncCallParser_negative(t *testing.T) {
	f := "alan(a b, c, 123, \"321\") xxx"
	res := FuncCallParser(f)
	if res != nil {
		t.Fatal("Func call should be not parsed")
	}
}
