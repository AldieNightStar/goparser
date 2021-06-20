package parser

import "testing"

func TestNext(t *testing.T) {
	txt := "Hello world!"

	if !IsNext(txt, "He") {
		t.Fail()
	}
	if !IsNext(txt[2:], "llo") {
		t.Fail()
	}
	if IsNext(txt, "x") {
		t.Fail()
	}

	if GetNext(txt, 4) != "Hell" {
		t.Fail()
	}
}

func TestUntil(t *testing.T) {
	txt := "Find the Temple"

	if Until(txt, " the") != "Find" {
		t.Fail()
	}
	if Until(txt, " Temple") != "Find the" {
		t.Fail()
	}
	if Until(txt[5:], " ") != "the" {
		t.Fail()
	}
}

func TestUntilOf(t *testing.T) {
	txt := "Kill Joe!EasterEgg@etc*"

	if s, s2 := UntilOf(txt, []string{"@", "*", "!"}); !(s == "Kill Joe" && s2 == "!") {
		t.Fail()
	}
	if s, s2 := UntilOf(txt[9:], []string{"*", "@"}); !(s == "EasterEgg" && s2 == "@") {
		t.Fail()
	}
}

func TestWhile(t *testing.T) {
	txt := "ababbcccaabb bbab"

	s := While(txt, "abc")

	if s != "ababbcccaabb" {
		t.Fail()
	}

	s = While(txt[12:], "abc")
	if s != "" {
		t.Fail()
	}
}
