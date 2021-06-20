package parser

import "strings"

func GetNext(text string, size int) string {
	if len(text) <= size {
		return text
	} else {
		return text[:size]
	}
}

func IsNext(text string, sub string) bool {
	t := GetNext(text, len(sub))
	return t == sub
}

func Until(text, sub string) string {
	ind := strings.Index(text, sub)
	if ind < 0 {
		return ""
	} else {
		return text[:ind]
	}
}

func UntilOf(text string, subs []string) (result string, substring string) {
	inds := make([]int, 0, len(subs))
	arr := make([]string, 0, len(subs))
	for i := 0; i < len(subs); i++ {
		ind := strings.Index(text, subs[i])
		if ind > -1 {
			inds = append(inds, ind)
			arr = append(arr, subs[i])
		}
	}
	low := 0xFFFF
	lows := ""
	n := 0
	for i := 0; i < len(inds); i++ {
		n = inds[i]
		if n < low {
			low = n
			lows = arr[i]
		}
	}
	if low == 0xFFFF {
		return "", ""
	}
	return text[:low], lows
}

func While(txt, set string) string {
	sb := &strings.Builder{}
	sb.Grow(32)
	for i := 0; i < len(txt); i++ {
		c := txt[i : i+1]
		if strings.Contains(set, c) {
			sb.WriteString(c)
		} else {
			break
		}
	}
	if sb.Len() < 1 {
		return ""
	}
	return sb.String()
}
