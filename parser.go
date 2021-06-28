package parser

type Result struct {
	// Token any object (Can be string or some struct)
	Token interface{}

	// Count of processed symbols
	Count int
}

type Parser func(string) *Result

func ParseOne(str string, parsers []Parser) *Result {
	for i := 0; i < len(parsers); i++ {
		res := parsers[i](str)
		if res != nil && res.Token != nil && res.Count > 0 {
			return res
		}
	}
	return nil
}

func Parse(str string, parsers []Parser) IteratorResult {
	cnt := 0
	return func() (result *Result, count int) {
		var err interface{} = nil
		defer func() {
			err = recover()
			if err != nil {
				result = nil
				count = cnt
			}
		}()
		if cnt >= len(str) {
			return nil, cnt
		}
		res := ParseOne(str[cnt:], parsers)
		if res != nil {
			cnt += res.Count
			return res, cnt
		} else {
			u := &UnknownToken{Value: str[0:1]}
			cnt += 1
			return &Result{Token: u, Count: 1}, cnt
		}
	}
}
