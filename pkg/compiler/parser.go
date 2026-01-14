package compiler

import (
	"fmt"
	"strings"
)

type NodeType int

const (
	NodeProgram NodeType = iota
	NodeCall
	NodeString
	NodeIdentifier
)

type Node struct {
	Type     NodeType
	Value    string
	Children []Node
}

type Parser struct {
	tokens  []Token
	current int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

func (p *Parser) Parse() Node {
	return Node{
		Type:     NodeProgram,
		Children: p.parseStatements(),
	}
}

func (p *Parser) parseStatements() []Node {
	var statements []Node

	for !p.isAtEnd() {
		stmt := p.parseStatement()
		if stmt.Type != NodeProgram || len(stmt.Children) > 0 {
			statements = append(statements, stmt)
		}
		p.skipNewlines()
	}

	return statements
}

func (p *Parser) parseStatement() Node {
	if p.match(TokenIdentifier) {
		ident := p.previous()
		return p.parseCall(ident)
	}

	return Node{Type: NodeProgram}
}

func (p *Parser) parseCall(ident Token) Node {
	call := Node{
		Type:  NodeCall,
		Value: ident.Value,
	}

	if p.match(TokenLeftParen) {
		if !p.check(TokenRightParen) {
			for {
				arg := p.parseExpression()
				call.Children = append(call.Children, arg)

				if p.check(TokenRightParen) {
					break
				}
			}
		}
		p.consume(TokenRightParen, "Expect ')' after arguments")
	}

	return call
}

func (p *Parser) parseExpression() Node {
	if p.match(TokenString) {
		return Node{
			Type:  NodeString,
			Value: strings.Trim(p.previous().Value, `"`),
		}
	}

	if p.match(TokenIdentifier) {
		return Node{
			Type:  NodeIdentifier,
			Value: p.previous().Value,
		}
	}

	return Node{Type: NodeProgram}
}

func (p *Parser) match(tokenType TokenType) bool {
	if p.check(tokenType) {
		p.advance()
		return true
	}
	return false
}

func (p *Parser) check(tokenType TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == TokenEOF
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}

func (p *Parser) consume(tokenType TokenType, message string) Token {
	if p.check(tokenType) {
		return p.advance()
	}
	panic(fmt.Sprintf("%s at line %d", message, p.peek().Line))
}

func (p *Parser) skipNewlines() {
	for p.match(TokenNewline) {
	}
}
