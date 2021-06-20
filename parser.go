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

func ParseChan(str string, parsers []Parser, c chan *Result) {
	cnt := 0
	for {
		if cnt >= len(str) {
			close(c)
			return
		}
		res := ParseOne(str[cnt:], parsers)
		if res != nil {
			c <- res
			cnt += res.Count
		} else {
			u := &UnknownToken{Value: str[0:1]}
			cnt += 1
			c <- &Result{Token: u, Count: 1}
		}
	}
}

func Parse(str string, parsers []Parser) []*Result {
	arr := make([]*Result, 0)
	c := make(chan *Result)
	go ParseChan(str, parsers, c)
	for res := range c {
		arr = append(arr, res)
	}
	return arr
}
