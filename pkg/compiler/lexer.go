package compiler

import (
	"fmt"
	"unicode"
)

type TokenType int

const (
	TokenEOF TokenType = iota
	TokenIdentifier
	TokenString
	TokenNumber
	TokenLeftParen
	TokenRightParen
	TokenColon
	TokenIndent
	TokenDedent
	TokenNewline
)

type Token struct {
	Type  TokenType
	Value string
	Line  int
}

type Lexer struct {
	input       string
	position    int
	line        int
	indentStack []int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:       input,
		position:    0,
		line:        1,
		indentStack: []int{0},
	}
}

func (l *Lexer) Tokenize() []Token {
	var tokens []Token

	for l.position < len(l.input) {
		ch := l.input[l.position]

		switch {
		case ch == '\n':
			tokens = append(tokens, Token{Type: TokenNewline, Line: l.line})
			l.position++
			l.line++
			l.handleIndent(&tokens)
		case unicode.IsSpace(rune(ch)) && ch != '\n':
			l.position++
		case ch == '(':
			tokens = append(tokens, Token{Type: TokenLeftParen, Value: "(", Line: l.line})
			l.position++
		case ch == ')':
			tokens = append(tokens, Token{Type: TokenRightParen, Value: ")", Line: l.line})
			l.position++
		case ch == ':':
			tokens = append(tokens, Token{Type: TokenColon, Value: ":", Line: l.line})
			l.position++
		case ch == '"':
			tokens = append(tokens, l.readString())
		case unicode.IsLetter(rune(ch)):
			tokens = append(tokens, l.readIdentifier())
		default:
			l.position++
		}
	}

	for len(l.indentStack) > 1 {
		l.indentStack = l.indentStack[:len(l.indentStack)-1]
		tokens = append(tokens, Token{Type: TokenDedent, Line: l.line})
	}

	tokens = append(tokens, Token{Type: TokenEOF, Line: l.line})
	return tokens
}

func (l *Lexer) handleIndent(tokens *[]Token) {
	indentLevel := 0
	for l.position < len(l.input) && (l.input[l.position] == ' ' || l.input[l.position] == '\t') {
		if l.input[l.position] == ' ' {
			indentLevel++
		}
		l.position++
	}

	currentIndent := l.indentStack[len(l.indentStack)-1]

	if indentLevel > currentIndent {
		l.indentStack = append(l.indentStack, indentLevel)
		*tokens = append(*tokens, Token{Type: TokenIndent, Line: l.line})
	} else if indentLevel < currentIndent {
		for len(l.indentStack) > 0 && l.indentStack[len(l.indentStack)-1] > indentLevel {
			l.indentStack = l.indentStack[:len(l.indentStack)-1]
			*tokens = append(*tokens, Token{Type: TokenDedent, Line: l.line})
		}
	}
}

func (l *Lexer) readString() Token {
	start := l.position
	l.position++

	for l.position < len(l.input) && l.input[l.position] != '"' {
		if l.input[l.position] == '\\' && l.position+1 < len(l.input) {
			l.position++
		}
		l.position++
	}

	l.position++
	value := l.input[start:l.position]
	return Token{Type: TokenString, Value: value, Line: l.line}
}

func (l *Lexer) readIdentifier() Token {
	start := l.position

	for l.position < len(l.input) && (unicode.IsLetter(rune(l.input[l.position])) || unicode.IsDigit(rune(l.input[l.position])) || l.input[l.position] == '_') {
		l.position++
	}

	value := l.input[start:l.position]
	return Token{Type: TokenIdentifier, Value: value, Line: l.line}
}

func (t Token) String() string {
	return fmt.Sprintf("Token{%s, %q, Line: %d}", t.typeName(), t.Value, t.Line)
}

func (t Token) typeName() string {
	switch t.Type {
	case TokenEOF:
		return "EOF"
	case TokenIdentifier:
		return "Identifier"
	case TokenString:
		return "String"
	case TokenNumber:
		return "Number"
	case TokenLeftParen:
		return "LeftParen"
	case TokenRightParen:
		return "RightParen"
	case TokenColon:
		return "Colon"
	case TokenIndent:
		return "Indent"
	case TokenDedent:
		return "Dedent"
	case TokenNewline:
		return "Newline"
	default:
		return "Unknown"
	}
}
