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

## Parsers out of box
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

## Pipes and filters
#### * Changer
* Collects the same sequence of tokens until name is changed
```go
// Create changer
changer := pipe.Changer()

// Now we can filter diff values
changer.Put("A", 12)
changer.Put("A", 24)
changer.Put("B", 144)
changer.Put("B", 188)
changer.Put("B", 3)
changer.Put("C", 13)

// And we are done
list := changer.Done()

// Result list will be
list[0] // A : 12, 24
list[1] // B : 144, 188, 3
list[2] // C : 13
```
