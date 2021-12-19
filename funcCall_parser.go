package parser

type FuncCallToken struct {
	FuncName string
	Args     []interface{}
}

func FuncCallParser(t string) *Result {
	// Parse first word as var
	pos := 0
	varRes := VariableParser(t)
	if varRes == nil {
		return nil
	}
	varName, varNameOk := varRes.Token.(*VariableToken)
	if !varNameOk {
		return nil
	}
	pos += varRes.Count
	// Parse '('
	if t[pos:pos+1] != "(" {
		return nil
	}
	pos += 1
	// Parse args
	token := &FuncCallToken{
		FuncName: varName.Name,
		Args:     make([]interface{}, 0, 8),
	}
	it := Parse(t[pos:], []Parser{
		FuncCallParser,
		StringParser,
		NumberParser,
		VariableParser,
		CreateSingleTokenParser(","), // Parse Separator as ','
		CreateSingleTokenParser(")"), // Parse ')' as end
	})

	commaArgCount := 1
	for {
		res, _ := it()
		if res == nil {
			break
		}
		// End with ')' symbol
		if singleTok, singleTokOk := res.Token.(*SingleToken); singleTokOk && singleTok.Name == ")" {
			pos += 1
			break
		}
		if commaTok, commaTokOk := res.Token.(*SingleToken); commaTokOk && commaTok.Name == "," {
			commaArgCount += 1
			pos += 1
			continue
		}
		if _, unknownTokenOk := res.Token.(*UnknownToken); unknownTokenOk {
			pos += res.Count
			continue
		}

		token.Args = append(token.Args, res.Token)
		if len(token.Args) != commaArgCount {
			return nil
		}
		pos += res.Count
	}
	return &Result{
		Token: token,
		Count: pos,
	}
}
