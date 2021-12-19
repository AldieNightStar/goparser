# Golang Parser - Lexer miniframework

# Import
```go
import parser "github.com/AldieNightStar/goparser"
```

## Usage
* Get all tokens array
```go
iter := Parse(text, parsers).ToArray()
```
* Get tokens iterator (We can break iteration)
```go
// text - text to parse
// parsers - array of 'Parser' objects
//
// returns: IteratorResults
iter := Parse(text, parsers)

// Pulls *Result one by one until `nil`
// Second return param is position - Symbol position
// Position can be used via Util_... functions to get line number at column number if needed
for {
	result, position := iter()
	if result == nil {
		break
	}
	// Do something with token and count
}

```
* Pull one token and calculate by ourself
```go
// text - text to parse
// parsers - array of 'Parser' objects
// returns: *Result
result := ParseOne(text, parsers)
```

## Parser example
* Parser it's a `func (string) *Result`
* Returns `*Result` with token and processed symbols count
* Can return `nil` when token is not supported
```go
func AbcParser(t string) *Result {
	s := parser.While(t, "abc")
	return &Result{s, len(s)}
}

// Then try to parse
arr := parser.Parse(text, []Parser{AbcParser}).ToArray()
```

## Result
```go
type Result struct {
	// Token any object (Can be string or some struct)
	Token interface{}

	// Count of processed symbols
	Count int
}
```

## Tools
```go
// Returns true when txt has "call" next
IsNext(txt, "call")

// Returns next 4 symbols from the text
GetNext(txt, 4)

// Reads text until some substring
// Returns "" if until-substring is not found
Until(txt, " end")

// Reads text until the closest one string from the list
// Returns text until string and stop-string itself as second param
// Returns "", "" if none of the elements are found
text, s := UntilOf(txt, string[]{"call", "end", "stop"})

// =======================
// Code line utils
// =======================

// Returns array of each line info:
//   From   - From which symbol current line is started
//   To     - Last line symbol position
//   Value  - Line itself
//   Number - Line number
//
// Better to use it once before the parsing.
lineInfos := Util_GetLines(sourceCode)

// Get specific line info by symbol position
// Best to know current line number etc
//
// lineInfos - []*LineInfo
// cnt       - int
oneLineInfo := Util_FindLineInfoAt(lineInfos, cnt)
```

## Result Iterator
* Iterator is not reusable
* To use iterator again - recreate it
```go
iter := Parse("some text ...", parsersArr)

// Pulls values one by one until `nil`
// Second return param is count of processed symbols (Symbol position)
token, count := iter()

// Get all tokens from the iterator as an Array
// Second return param is count of processed symbols (Symbol position)
allTokensArr, count := iter.ToArray()

// Get tokens which is ok with your function criterias.
// To allow token to be in filtered list - func should return true
// ! - Skips unwanted tokens
// Second return param is count of processed symbols (Symbol position)
filteredTokensArr, count := iter.FilterArray(func (res *Result) bool {
	// return true | false
})

// Returns list of tokens until some special token
// - Last token is not included
// Second return param is count of processed symbols (Symbol position)
untilSomeTokenArr, count := iter.UntilArray(func (res *Result) bool {
	// return true | false
	//	true  - this is until-token (end-token)
	//  false - this is not until-token (end-token)
})

// Retuns up to {n} tokens or less (depends on count)
// Second return param is count of processed symbols (Symbol position)
fewTokensArr, count := iter.FewArray(5)
```

## Parsers out of the box
```go
// Parses strings:
// "string of text", 'string of text', `string of text`
// Escaping works with `\` symbol. Also parses: \n \t \0 \r
//
// Returns:   *StringToken(Value, Quote)
res := StringParser(text)

// Parses numbers of float64
// Supports dot values. 1.32, 4.678 etc. But not two dots
//
// Returns:   *NumberToken(Value: float64)
res := NumberParser(text)

// Parses variable string (These which are used in simple languages)
// 	Samples:
//		ProfileName, name, second_name, prof1, prof2, etc
//
// Returns:   *VariableToken(Name: string)
res := VariableParser(text)

// Parse function call
//  Samples:
//		Func(a, b, c)
//      Func(GetX(), GetY(), GetZ()) 
//
// Returns: 	*FuncCallToken(FuncName: string, Args: []inteface{})
res := FuncCallParser(text)
```

## Parser creators out of the box
```go
// Create parser which parses only ONE hardcoded peace of text and creates a token
// Can be used as an end of recursive parsing
// Example:
// 		endParser := SingleToken("}")
// Then can be used:
//		ParseOne(text, []Parser{parser1, parser2, endParser})
//
// Returns:   *SingleToken(Name: string)
endParser := CreateSingleTokenParser("end")
```