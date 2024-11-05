package lexer

import "json-parser-and-query-tool/pkg/token"

// to perform lexical analysis
type Lexer struct {
	Input        []rune
	char         rune // current character under examination
	position     int  // current position in the input (points to current char)
	readPosition int  // current read position in the input (after current read char)
	line         int  // Line number for better error reporting
}

// Creates and returns pointer to the lexer
func New(input string) *Lexer {
	l := &Lexer{Input: []rune(input)}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.Input) {
		// we have exceeded the string, EOF or empty or haven't read anything thing
		// 0 is the ASCII equivalent of NULL character
		l.char = 0
	} else {
		l.char = l.Input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		if l.char == '\n' {
			l.line++
		}
		l.readChar()
	}
}

func newToken(tokenType token.Type, line, start, end int, char ...rune) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(char),
		Line:    line,
		Start:   start,
		End:     end,
	}
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhitespace()

	switch l.char {
	case '{':
		t = newToken(token.LeftBrace, l.line, l.position, l.position+1, l.char)
	case '}':
		t = newToken(token.RightBrace, l.line, l.position, l.position+1, l.char)
	case '[':
		t = newToken(token.LeftBracket, l.line, l.position, l.position+1, l.char)
	case ']':
		t = newToken(token.RightBracket, l.line, l.position, l.position+1, l.char)
	case ':':
		t = newToken(token.Colon, l.line, l.position, l.position+1, l.char)
	case ',':
		t = newToken(token.Comma, l.line, l.position, l.position+1, l.char)
	case '"':
		t.Type = token.String
		t.Literal = l.readString()
		t.Line = l.line
		t.Start = l.position
		t.End = l.position + 1
	case 0:
		t.Type = token.EOF
		t.Literal = ""
		t.Line = l.line
	default:
		if isLetter(l.char) {
			t.Start = l.position
			ident := l.readIdentifier(isLetter)
			t.Literal = ident
			t.Line = l.line

			tokenType, err := token.LookupIdentifier(ident)
			if err != nil {
				t.Type = token.Illegal
				return t
			}
			t.Type = tokenType
			t.End = l.position
			return t
		} else if isNumber(l.char) {
			t.Start = l.position
			ident := l.readIdentifier(isNumber)
			t.Literal = ident
			t.Line = l.line
			t.Type = token.Number
			t.End = l.position
			return t
		}
		t = newToken(token.Illegal, l.line, 1, 2, l.char)
	}

	l.readChar()

	return t
}

func (l *Lexer) readString() string {
	position := l.position + 1 // since l.position is "

	for {
		prevChar := l.char
		l.readChar()
		if (l.char == '"' && prevChar != '\\') || l.char == 0 { // unescaped or EOF
			break
		}
	}
	return string(l.Input[position:l.position])
}

type identifierFunc func(char rune) bool

func isNumber(char rune) bool {
	if (char >= '0' && char <= '9') || char == '.' || char == '-' {
		return true
	}
	return false
}

func isLetter(char rune) bool {
	if char >= 'a' && char <= 'z' {
		return true
	}
	return false
}

func (l *Lexer) readIdentifier(f identifierFunc) string {
	position := l.position

	for f(l.char) {
		l.readChar()
	}

	return string(l.Input[position:l.position])
}
