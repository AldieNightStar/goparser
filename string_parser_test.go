package parser

import (
	"testing"
)

func TestStringParser(t *testing.T) {
	txt := "'abc def' \"killer\" `123\\`321`"

	res := StringParser(txt)
	if res == nil || res.Token == nil {
		t.Fatal("Result is null or token is null")
	}
	if res.Count != 9 {
		t.Fatal("Count is not 9. ", res.Count)
	}

	if st, ok := res.Token.(StringToken); ok {
		if st.Value != "abc def" {
			t.Fatal("Value for string token is not valid.", st.Value)
		}
		if st.Quote != '\'' {
			t.Fatal("Value of qoute is not valid.", st.Quote)
		}
	}

	res = StringParser(txt[19:])
	if res == nil {
		t.Fatal("Result is null")
	}

	if st, ok := res.Token.(StringToken); ok {
		if st.Value != "abc`def" {
			t.Fatal("Value for string token is not valid.", st.Value)
		}
		if st.Quote != '\'' {
			t.Fatal("Value of qoute is not valid.", st.Quote)
		}
	}
}
