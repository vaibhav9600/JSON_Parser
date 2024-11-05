package token

import "fmt"

type Type string

const (
	Illegal Type = "ILLEGAL"
	// End of file
	EOF Type = "EOF"

	// Literals
	String Type = "STRING"
	Number Type = "NUMBER"

	// The six structural tokens
	LeftBrace    Type = "{"
	RightBrace   Type = "}"
	LeftBracket  Type = "["
	RightBracket Type = "]"
	Comma        Type = ","
	Colon        Type = ":"

	// Values
	True  Type = "TRUE"
	False Type = "FALSE"
	Null  Type = "NULL"
)

type Token struct {
	Type    Type
	Literal string
	Line    int
	Start   int
	End     int
}

var validJSONTypes = map[string]Type{
	"true":  True,
	"false": False,
	"null":  Null,
}

func LookupIdentifier(identifier string) (Type, error) {
	if token, exists := validJSONTypes[identifier]; exists {
		return token, nil
	}

	return "", fmt.Errorf("expected a valid JSON identifier. Found: %s", identifier)
}

