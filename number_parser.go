package parser

import (
	"strconv"
	"strings"
)

type NumberToken struct {
	Value float64
}

func NumberParser(t string) *Result {
	sb := strings.Builder{}
	dot := false
	for i := 0; i < len(t); i++ {
		c := t[i]
		if c == '.' {
			if !dot {
				dot = true
				sb.WriteByte(c)
				continue
			} else {
				break
			}
		} else if strings.Contains("0123456789", string(c)) {
			sb.WriteByte(c)
			continue
		} else {
			break
		}
	}
	if sb.Len() < 1 {
		return nil
	}
	str := sb.String()
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return nil
	}
	return &Result{&NumberToken{num}, len(str)}
}
