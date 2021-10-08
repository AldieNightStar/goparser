package parser

import "testing"

func TestCreateSingleTokenParser(t *testing.T) {
	parser := CreateSingleTokenParser("end")
	if parser("remaster") != nil {
		t.Fatal("Wrong token. 'remaster' isn't 'end' token")
	}

	res := parser("end of life")
	if token, ok := res.Token.(*SingleToken); ok {
		if token.Name != "end" {
			t.Fatal("'end' is not 'end'?")
		}
	}
}
