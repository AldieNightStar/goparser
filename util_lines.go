package parser

import "strings"

type LineInfo struct {
	Value  string
	From   int
	To     int
	Number int
}

func Util_GetLines(str string) []*LineInfo {
	list := make([]*LineInfo, 0, 32)
	size := len(str)
	sb := __str_builder()
	lastFrom := 0
	num := 0
	for i := 0; i < size; i++ {
		c := str[i]
		if c == '\n' {
			l := &LineInfo{
				Value:  sb.String(),
				From:   lastFrom,
				To:     i,
				Number: num,
			}
			sb = __str_builder()
			lastFrom = i + 1
			num += 1
			list = append(list, l)
			continue
		}
		sb.WriteByte(c)
	}
	if sb.Len() > 0 {
		l := &LineInfo{
			Value:  sb.String(),
			From:   lastFrom,
			To:     len(str),
			Number: num,
		}
		list = append(list, l)
	}
	return list
}

func __str_builder() *strings.Builder {
	sb := &strings.Builder{}
	sb.Grow(128)
	return sb
}

func Util_FindLineInfoAt(lines []*LineInfo, cnt int) *LineInfo {
	for _, line := range lines {
		if line.From <= cnt && line.To >= cnt {
			return line
		}
	}
	return nil
}
