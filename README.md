# Golang Parser - Lexer miniframework

# Import
```go
import parser "github.com/AldieNightStar/goparser"
```

## Usage
```go
// text - text to parse
// parsers - array of 'Parser' objects
// returns: []*Result
results := Parse(text, parsers)

// text - text to parse
// parsers - array of 'Parser' objects
// returns: *Result
result := ParseOne(text, parsers)

// text - text to parse
// parsers - array of 'Parser' objects
// ch - channel of '*Result' objects (Will be closed after parsing)
go ParseChan(text, parsers, ch)
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
arr := parser.Parse(text, &AbcParser)
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

// Returns next 4 symbols from text
GetNext(txt, 4)

// Reads text until some substring
// Returns "" if until-substring is not found
Until(txt, " end")

// Reads text until the closest one string from list
// Returns text until string and stop-string itself as second param
// Returns "", "" if none of the elements are found
text, s :=  UntilOf(txt, string[]{"call", "end", "stop"})
```