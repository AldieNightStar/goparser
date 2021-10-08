package parser

import "testing"

func TestVariableParser(t *testing.T) {
	res := VariableParser("the_variadic blind_document")

	if tok, ok := res.Token.(*VariableToken); ok {
		if tok.Name != "the_variadic" {
			t.Fatal("Variable name is not: the_variadic")
		}
	}

	res = VariableParser("1Dog")

	if res != nil {
		t.Fatal("1Dog cannot be variable. Should be nil")
	}

	res = VariableParser("Hero]Until]Grave")
	if tok, ok := res.Token.(*VariableToken); ok {
		if tok.Name != "Hero" {
			t.Fatal("Variable name is not: Hero")
		}
		if res.Count != 4 {
			t.Fatal("Size should be 4: Hero")
		}
	}
}
