package parser

import (
	"regexp"
	"strings"
)

type VariableToken struct {
	Name string
}

var variableRegex, _ = regexp.Compile(`([a-z]|[A-Z]|[0-9]|\_)*`)

func VariableParser(t string) *Result {
	name := variableRegex.FindString(t)
	if len(name) < 1 {
		return nil
	}
	if strings.Contains("0123456789", name[0:1]) {
		return nil
	}
	return &Result{
		Token: &VariableToken{
			Name: name,
		},
		Count: len(name),
	}
}
