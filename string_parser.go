package parser

import "strings"

type StringToken struct {
	Value string
	Quote byte
}

func StringParser(s string) *Result {
	if len(s) < 2 {
		return nil
	}
	size := len(s)
	sb := &strings.Builder{}
	var quote byte = 0
	escape := false
	escapes := 0
	lastIsQuote := false
	for i := 0; i < size; i++ {
		if i == 0 {
			quote = s[i]
			if !(quote == '"' || quote == '\'' || quote == '`') {
				return nil
			}
			continue
		}
		c := s[i]
		if escape {
			if c == 'n' {
				c = '\n'
			} else if c == 't' {
				c = '\t'
			} else if c == '0' {
				c = 0
			} else if c == 'r' {
				c = '\r'
			}
			escapes += 1
			sb.WriteByte(c)
			escape = false
			continue
		}
		if c == '\\' {
			escape = true
			continue
		}
		if c == quote {
			lastIsQuote = true
			break
		}
		sb.WriteByte(c)
	}
	if !lastIsQuote {
		return nil
	}
	sbLen := sb.Len()
	if sbLen < 1 {
		return &Result{&StringToken{"", quote}, 2}
	}
	return &Result{&StringToken{sb.String(), quote}, sbLen + escapes + 2}
}
