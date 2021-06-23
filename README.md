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
for {
	result := iter()
	if result == nil {
		break
	}
	// Do something with tokens
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
arr := parser.Parse(text, &AbcParser).ToArray()
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
text, s :=  UntilOf(txt, string[]{"call", "end", "stop"})
```

## Result Iterator
* Iterator is not reusable
* To use iterator again - recreate it
```go
iter := Parse("some text ...", parsersArr)

// Get all tokens from the iterator as an Array
allTokensArr := iter.ToArray()

// Get tokens which is ok with your function criterias.
// To allow token to be in filtered list - func should return true
// ! - Skips unwanted tokens
filteredTokensArr := iter.FilterArray(func (res *Result) bool {
	// return true | false
})

// Returns list of tokens until some special token
// - Last token is not included
untilSomeTokenArr := iter.UntilArray(func (res *Result) bool {
	// return true | false
	//	true  - this is until-token (end-token)
	//  false - this is not until-token (end-token)
})

// Retuns up to {n} tokens or less (depends on count)
fewTokensArr := iter.FewArray(5)
```

## Parsers out of the box
```go
// Parses strings:
// "string of text", 'string of text', `string of text`
// Escaping works with `\` symbol. Also parses: \n \t \0 \r
//
// Returns: StringToken(Value, Quote)
res := StringParser(text)

// Parses numbers of float64
// Supports dot values. 1.32, 4.678 etc. But not two dots
//
// Returns: NumberToken(Value: float64)
res := NumberParser(text)
```